package models

import "time"

type Todo struct {
	ID        int `gorm:"primaryKey"`
	Task      string
	Done      bool
	Created   time.Time
	Completed time.Time
}
