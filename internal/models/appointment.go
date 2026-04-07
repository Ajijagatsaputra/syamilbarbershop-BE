package models

import (
	"database/sql/driver"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type AppointmentStatus string

const (
	StatusPending   AppointmentStatus = "pending"
	StatusConfirmed AppointmentStatus = "confirmed"
	StatusCompleted AppointmentStatus = "completed"
	StatusCancelled AppointmentStatus = "cancelled"
)

type JSONDate time.Time

func (j *JSONDate) UnmarshalJSON(b []byte) error {
	s := string(b)
	if s == "null" {
		return nil
	}
	t, err := time.Parse("\"2006-01-02\"", s)
	if err != nil {
		return err
	}
	*j = JSONDate(t)
	return nil
}

func (j JSONDate) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", time.Time(j).Format("2006-01-02"))), nil
}

func (j JSONDate) Value() (driver.Value, error) {
	return time.Time(j), nil
}

func (j *JSONDate) Scan(value interface{}) error {
	t, ok := value.(time.Time)
	if !ok {
		return fmt.Errorf("failed to scan JSONDate")
	}
	*j = JSONDate(t)
	return nil
}

type Appointment struct {
	ID         uint              `gorm:"primaryKey" json:"id"`
	CustomerID uint              `json:"customer_id"`
	Customer   Customer          `gorm:"foreignKey:CustomerID" json:"customer"`
	ServiceID  uint              `json:"service_id"`
	Service    Service           `gorm:"foreignKey:ServiceID" json:"service"`
	Date       JSONDate          `gorm:"type:date;not null" json:"date"`
	Time       string            `gorm:"type:varchar(10);not null" json:"time"` // e.g., "10:00"
	Status        AppointmentStatus `gorm:"type:varchar(20);default:'pending'" json:"status"`
	PaymentMethod string            `gorm:"type:varchar(50)" json:"payment_method"`
	Price         float64           `gorm:"type:decimal(10,2)" json:"price"`
	CreatedAt  time.Time         `json:"createdAt"`
	UpdatedAt  time.Time         `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt    `gorm:"index" json:"-"`
}
