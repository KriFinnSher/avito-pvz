package reception

import (
	"avito-pvz/internal/models"
	"context"
	"github.com/google/uuid"
)

//go:generate mockery --name=Repository --output=../../mocks/reception --with-expecter

// Repository defines methods for managing reception records tied to PVZ locations
type Repository interface {
	Create(ctx context.Context, reception models.Reception) error
	CloseLast(ctx context.Context, pvzID uuid.UUID) error
	IsOpen(ctx context.Context, receptionID uuid.UUID) bool
	GetLast(ctx context.Context, pvzID uuid.UUID) (models.Reception, error)
	GetAllForPVZ(ctx context.Context, pvzID uuid.UUID) ([]models.Reception, error)
}
