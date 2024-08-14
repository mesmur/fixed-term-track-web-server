package cron

import (
	"time"

	"github.com/MESMUR/fixed-term-track-web-server/internal/clients"
	"github.com/MESMUR/fixed-term-track-web-server/internal/models"
	"github.com/MESMUR/fixed-term-track-web-server/pkg/logger"
	"github.com/MESMUR/fixed-term-track-web-server/repositories"
)

// EventReader - is responsible for reading events from the database and processing them
type EventReader interface {
	CheckEvents()
}

type eventReader struct {
	eventRepo repositories.EventRepository
	telegram  clients.TelegramSdk
}

func NewEventReader(eventRepo repositories.EventRepository, telegram clients.TelegramSdk) EventReader {
	return &eventReader{eventRepo: eventRepo, telegram: telegram}
}

func (er *eventReader) CheckEvents() {
	ticker := time.NewTicker(8 * time.Hour)
	defer ticker.Stop()

	for range ticker.C {
		logger.Log.Info("Checking for scheduled events")

		events, err := er.eventRepo.FindScheduledEvents()

		if err != nil {
			logger.Sugar.Error("Error fetching scheduled events: ", err)
			continue
		}

		for _, event := range events {
			logger.Sugar.Infof("Processing event: %v", event)

			er.processEvent(event)
		}
	}
}

func (er *eventReader) processEvent(event models.Event) {
	switch event.EventType {
	case "MONTHLY_RETURN_NOTIFICATION":
	case "MATURITY_RETURN_NOTIFICATION":
		er.handleReturnNotification(event)
	default:
		logger.Sugar.Errorf("Unknown event type: %s", event.EventType)
	}
}

func (er *eventReader) handleReturnNotification(event models.Event) {
	err := er.telegram.SendMessage(event.Message)

	if err != nil {
		logger.Sugar.Errorf("Error sending telegram message for event: %d.\n%+v", event.ID, err)
	}

	err = er.eventRepo.UpdateStatus(event.ID)

	if err != nil {
		logger.Sugar.Errorf("Error updating event status for event: %d.\n%+v", event.ID, err)
	}
}
