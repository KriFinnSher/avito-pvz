package pvz

import (
	"avito-pvz/internal/models"
	"context"
)

// Repository defines methods for managing PVZ (pickup point) records
type Repository interface {
	Create(ctx context.Context, pvz models.PVZ) error
	GetAll(ctx context.Context) ([]models.PVZ, error)
}
