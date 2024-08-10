package cron

import (
	"github.com/MESMUR/fixed-term-track-web-server/internal/clients"
	"github.com/MESMUR/fixed-term-track-web-server/pkg/logger"
	"github.com/MESMUR/fixed-term-track-web-server/repositories"
	"time"
)

type EventReader struct {
	eventRepo repositories.EventRepository
	telegram  clients.TelegramSdk
}

func NewEventReader(eventRepo repositories.EventRepository, telegram clients.TelegramSdk) *EventReader {
	return &EventReader{eventRepo: eventRepo, telegram: telegram}
}

func (er *EventReader) CheckEvents() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			logger.Log.Info("Checking for scheduled events")

			err := er.telegram.SendTelegramMessage("Checking for scheduled events")

			if err != nil {
				logger.Sugar.Error("Error sending telegram message: ", err)
			}

			events, err := er.eventRepo.FindScheduledEvents()

			if err != nil {
				logger.Sugar.Error("Error fetching scheduled events: ", err)
				continue
			}

			for _, event := range events {
				// Process each event
				logger.Sugar.Infof("Processing event: %v", event)
				// Add your event processing logic here

				// Update the event status
				err := er.eventRepo.UpdateStatus(event.ID)

				if err != nil {
					logger.Sugar.Error("Error updating event status: ", err)
				}
			}
		}
	}
}
