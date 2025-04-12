package handlers

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	Role  UserRole  `json:"role"`
}

type PVZ struct {
	ID               uuid.UUID `json:"id"`
	RegistrationDate time.Time `json:"registrationDate"`
	City             PvzCity   `json:"city"`
}

type Reception struct {
	ID       uuid.UUID       `json:"id"`
	DateTime time.Time       `json:"dateTime"`
	PvzID    uuid.UUID       `json:"pvzId"`
	Status   ReceptionStatus `json:"status"`
}

type Product struct {
	ID          uuid.UUID   `json:"id"`
	DateTime    time.Time   `json:"dateTime"`
	Type        ProductType `json:"type"`
	ReceptionID uuid.UUID   `json:"receptionId"`
}

type UserRole string

type PvzCity string

type ReceptionStatus string

type ProductType string

const (
	EmployeeRole  UserRole = "employee"
	ModeratorRole UserRole = "moderator"
)

const (
	KazanCity  PvzCity = "Казань"
	MoscowCity PvzCity = "Москва"
	SaintCity  PvzCity = "Санкт-Петербург"
)

const (
	InProgressStatus ReceptionStatus = "in_progress"
	CloseStatus      ReceptionStatus = "close"
)

const (
	ElectronicType ProductType = "электроника"
	ClothesType    ProductType = "одежда"
	ShoesType      ProductType = "обувь"
)

type DummyRequest struct {
	Role UserRole `json:"role"`
}

type DummyResponse struct {
	Token string `json:"Token"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type RegisterRequest struct {
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Role     UserRole `json:"role"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ReceptionRequest struct {
	PvzID uuid.UUID `json:"pvzId"`
}

type ProductRequest struct {
	Type  ProductType `json:"type"`
	PvzID uuid.UUID   `json:"pvzId"`
}

type ReceptionWithProducts struct {
	Reception Reception `json:"reception"`
	Products  []Product `json:"products"`
}

type PVZResponse struct {
	PVZ        PVZ                     `json:"pvz"`
	Receptions []ReceptionWithProducts `json:"receptions"`
}

type PVZPaginationResponse struct {
	TotalCount int           `json:"totalCount"`
	Page       int           `json:"page"`
	Limit      int           `json:"limit"`
	PVZs       []PVZResponse `json:"pvzs"`
}

type GetPVZQueryParams struct {
	StartDate string `json:"startDate,omitempty"`
	EndDate   string `json:"endDate,omitempty"`
	Page      int    `json:"page,omitempty"`
	Limit     int    `json:"limit,omitempty"`
}

const (
	TestUser1Email = "testuser1@example.com"
	TestUser2Email = "testuser2@example.com"
	TestUserPass   = "password123"
)
