package models

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ID          uuid.UUID `json:"id" db:"id"`
	DateTime    time.Time `json:"dateTime" db:"date_time"`
	Type        string    `json:"type" db:"type"`
	ReceptionId uuid.UUID `json:"receptionId" db:"reception_id"`
}
