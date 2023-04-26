package models

import (
	"gorm.io/gorm"
	"time"
)

type Post struct {
	ID        uint `gorm:"primarykey"`
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	ProcessingResult *ProcessingResult `json:"processing_result,omitempty"`
}

type ProcessingResult struct {
	ID              uint `gorm:"primarykey"`
	PostID          uint
	NumberOfSymbols uint
	NumberOfWords   uint
	ReadingTime     uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
