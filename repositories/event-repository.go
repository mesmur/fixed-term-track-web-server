package repositories

import (
	"github.com/MESMUR/fixed-term-track-web-server/internal/database/models"
	"gorm.io/gorm"
	"time"
)

type EventRepository interface {
	FindByID(eventId uint) (*models.Event, error)
	Create(event *models.Event) error
	FindScheduledEvents() ([]models.Event, error)
	UpdateStatus(eventId uint) error
}

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventRepository{db}
}

func (r *eventRepository) FindByID(eventId uint) (*models.Event, error) {
	var event models.Event
	err := r.db.Where("id = ?", eventId).First(&event).Error

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (r *eventRepository) Create(event *models.Event) error {
	if err := r.db.Create(event).Error; err != nil {
		return err
	}

	return nil
}

func (r *eventRepository) FindScheduledEvents() ([]models.Event, error) {
	var events []models.Event
	err := r.db.Where("scheduled_time <= ? AND status = ?", time.Now(), "SCHEDULED").Find(&events).Error
	return events, err
}

func (r *eventRepository) UpdateStatus(eventId uint) error {
	event, err := r.FindByID(eventId)

	if err != nil {
		return err
	}

	event.Status = "COMPLETED"
	err = r.db.Save(event).Error

	return err
}
