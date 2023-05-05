package models

import "time"

type Worker struct {
	ID         uint `gorm:"primary_key"`
	Name       string
	Hostname   string
	WorkerType string

	StartedAt    time.Time
	LastActiveAt time.Time
}
