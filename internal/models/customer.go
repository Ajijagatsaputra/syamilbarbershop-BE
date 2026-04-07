package models

import (
	"time"
	"gorm.io/gorm"
)

type Customer struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"type:varchar(100);not null" json:"name"`
	Phone     string         `gorm:"type:varchar(20);not null" json:"phone"`
	Email     string         `gorm:"type:varchar(100)" json:"email"`
	LastVisit *time.Time      `json:"lastVisit"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
