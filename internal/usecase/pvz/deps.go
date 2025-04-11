package pvz

import (
	"avito-pvz/internal/models"
	"context"
)

//go:generate mockery --name=Repository --output=../../mocks/pvz --with-expecter --case=underscore

// Repository defines methods for managing PVZ (pickup point) records
type Repository interface {
	Create(ctx context.Context, pvz models.PVZ) error
	GetAll(ctx context.Context) ([]models.PVZ, error)
}
