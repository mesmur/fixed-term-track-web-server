package models

import "time"

type Event struct {
	ID            uint `gorm:"primaryKey"`
	ResourceID    uint // ID of the resource that was responsible for this event being scheduled
	EventType     string
	ScheduledTime time.Time
	Status        string
}
