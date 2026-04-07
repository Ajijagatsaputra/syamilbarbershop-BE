package models

import (
	"time"
	"gorm.io/gorm"
)

type ServiceCategory string

const (
	CategoryHaircut  ServiceCategory = "Haircut"
	CategoryGrooming ServiceCategory = "Grooming"
	CategoryStyling  ServiceCategory = "Styling"
)

type ServiceStatus string

const (
	StatusActive   ServiceStatus = "active"
	StatusInactive ServiceStatus = "inactive"
)

type Service struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"type:varchar(100);not null" json:"name"`
	Price     float64        `gorm:"type:decimal(10,2);not null" json:"price"`
	Duration  int            `gorm:"not null" json:"duration"` // in minutes
	Category  ServiceCategory `gorm:"type:varchar(20);not null" json:"category"`
	Status    ServiceStatus   `gorm:"type:varchar(20);default:'active'" json:"status"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
