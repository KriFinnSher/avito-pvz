package reception

import (
	"avito-pvz/internal/models"
	"context"
	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, reception models.Reception) error
	CloseLast(ctx context.Context, pvzID uuid.UUID) error
	IsOpen(ctx context.Context, receptionID uuid.UUID) bool
	GetLast(ctx context.Context, pvzID uuid.UUID) (models.Reception, error)
	GetAllForPVZ(ctx context.Context, pvzID uuid.UUID) ([]models.Reception, error)
}
