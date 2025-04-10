package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID    uuid.UUID `json:"id" db:"id"`
	Email string    `json:"email" db:"email"`
	Role  string    `json:"role" db:"role"`
	Hash  string    `json:"hash" db:"hash"`
}
