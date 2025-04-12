package usecase

import (
	"avito-pvz/internal/models"
	"avito-pvz/internal/usecase/user"
	"context"
	"log/slog"
)

// UserUseCase contains business logic for managing user-related operations
type UserUseCase struct {
	userRepo user.Repository
	logger   *slog.Logger
}

// NewUserUseCase creates a new UserUseCase with the given user repository and logger
func NewUserUseCase(uRepo user.Repository, logger *slog.Logger) *UserUseCase {
	return &UserUseCase{
		userRepo: uRepo,
		logger:   logger,
	}
}

// RegisterUser registers a new user in the system
func (u *UserUseCase) RegisterUser(ctx context.Context, user models.User) error {
	u.logger.Info("Attempting to register user", "email", user.Email)

	err := u.userRepo.Create(ctx, user)
	if err != nil {
		u.logger.Error("Failed to register user", "email", user.Email, "error", err)
		return err
	}

	u.logger.Info("User registered successfully", "email", user.Email)
	return nil
}

// GetUserByEmail retrieves a user by their email address
func (u *UserUseCase) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	u.logger.Info("Attempting to get user by email", "email", email)

	myUser, err := u.userRepo.GetByEmail(ctx, email)
	if err != nil {
		u.logger.Error("Failed to get user by email", "email", email, "error", err)
		return models.User{}, err
	}

	u.logger.Info("User fetched successfully", "email", email)
	return myUser, nil
}

// CheckUserExists checks if a user with the given email exists in the system
func (u *UserUseCase) CheckUserExists(ctx context.Context, email string) (bool, error) {
	u.logger.Info("Checking if user exists", "email", email)

	exists, err := u.userRepo.Exists(ctx, email)
	if err != nil {
		u.logger.Error("Failed to check if user exists", "email", email, "error", err)
		return false, err
	}

	u.logger.Info("User existence checked", "email", email, "exists", exists)
	return exists, nil
}
