package user

import (
	"avito-pvz/internal/models"
	"context"
)

type Repository interface {
	Create(ctx context.Context, user models.User) error
	GetByEmail(ctx context.Context, email string) (models.User, error)
	Exists(ctx context.Context, email string) (bool, error)
}
