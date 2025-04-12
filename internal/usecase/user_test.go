package usecase

import (
	"avito-pvz/internal/models"
	mockUser "avito-pvz/internal/repository/mocks/user"
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"os"
	"testing"
)

func TestRegisterUser_Success(t *testing.T) {
	ctx := context.Background()
	userEmail := "test@example.com"

	userRepo := mockUser.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	newUser := models.User{
		ID:    uuid.New(),
		Email: userEmail,
		Role:  "user",
		Hash:  "hashed_password",
	}

	userRepo.On("Create", ctx, newUser).Return(nil)

	useCase := NewUserUseCase(userRepo, logger)

	err := useCase.RegisterUser(ctx, newUser)

	assert.NoError(t, err)
	userRepo.AssertCalled(t, "Create", ctx, newUser)
}

func TestRegisterUser_Error(t *testing.T) {
	ctx := context.Background()
	userEmail := "test@example.com"

	userRepo := mockUser.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	newUser := models.User{
		ID:    uuid.New(),
		Email: userEmail,
		Role:  "user",
		Hash:  "hashed_password",
	}

	userRepo.On("Create", ctx, newUser).Return(errors.New("failed to create user"))

	useCase := NewUserUseCase(userRepo, logger)

	err := useCase.RegisterUser(ctx, newUser)

	assert.Error(t, err)
	userRepo.AssertCalled(t, "Create", ctx, newUser)
}

func TestGetUserByEmail_Success(t *testing.T) {
	ctx := context.Background()
	userEmail := "test@example.com"

	userRepo := mockUser.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	expectedUser := models.User{
		ID:    uuid.New(),
		Email: userEmail,
		Role:  "user",
		Hash:  "hashed_password",
	}

	userRepo.On("GetByEmail", ctx, userEmail).Return(expectedUser, nil)

	useCase := NewUserUseCase(userRepo, logger)

	user, err := useCase.GetUserByEmail(ctx, userEmail)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
	userRepo.AssertCalled(t, "GetByEmail", ctx, userEmail)
}

func TestGetUserByEmail_Error(t *testing.T) {
	ctx := context.Background()
	userEmail := "test@example.com"

	userRepo := mockUser.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	userRepo.On("GetByEmail", ctx, userEmail).Return(models.User{}, errors.New("user not found"))

	useCase := NewUserUseCase(userRepo, logger)

	user, err := useCase.GetUserByEmail(ctx, userEmail)

	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())
	assert.Equal(t, models.User{}, user)
	userRepo.AssertCalled(t, "GetByEmail", ctx, userEmail)
}

func TestCheckUserExists_Success(t *testing.T) {
	ctx := context.Background()
	userEmail := "test@example.com"

	userRepo := mockUser.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	userRepo.On("Exists", ctx, userEmail).Return(true, nil)

	useCase := NewUserUseCase(userRepo, logger)

	exists, err := useCase.CheckUserExists(ctx, userEmail)

	assert.NoError(t, err)
	assert.True(t, exists)
	userRepo.AssertCalled(t, "Exists", ctx, userEmail)
}

func TestCheckUserExists_Error(t *testing.T) {
	ctx := context.Background()
	userEmail := "test@example.com"

	userRepo := mockUser.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	userRepo.On("Exists", ctx, userEmail).Return(false, errors.New("error checking user existence"))

	useCase := NewUserUseCase(userRepo, logger)

	exists, err := useCase.CheckUserExists(ctx, userEmail)

	assert.Error(t, err)
	assert.False(t, exists)
	userRepo.AssertCalled(t, "Exists", ctx, userEmail)
}
