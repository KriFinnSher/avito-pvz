package product

import (
	"avito-pvz/internal/models"
	"context"
	"github.com/google/uuid"
)

// Repository defines methods for managing product records related to receptions
type Repository interface {
	DeleteLast(ctx context.Context, receptionID uuid.UUID) error
	AddOne(ctx context.Context, product models.Product) error
	GetForReception(ctx context.Context, receptionID uuid.UUID) ([]models.Product, error)
}
