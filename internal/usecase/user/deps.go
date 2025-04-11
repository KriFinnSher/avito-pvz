package user

import (
	"avito-pvz/internal/models"
	"context"
)

//go:generate mockery --name=Repository --output=../../mocks/user --with-expecter --case=underscore

// Repository defines methods for managing users
type Repository interface {
	Create(ctx context.Context, user models.User) error
	GetByEmail(ctx context.Context, email string) (models.User, error)
	Exists(ctx context.Context, email string) (bool, error)
}
