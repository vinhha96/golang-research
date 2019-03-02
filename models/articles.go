package models

import (
	"time"
)

type Article struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	Title   string `json:"title"`
	Content string `json:"content"`
}
