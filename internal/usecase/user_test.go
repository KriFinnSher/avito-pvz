package usecase

import (
	mockUser "avito-pvz/internal/mocks/user"
	"avito-pvz/internal/models"
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

	// Моки
	userRepo := mockUser.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Пример пользователя
	newUser := models.User{
		ID:    uuid.New(),
		Email: userEmail,
		Role:  "user",
		Hash:  "hashed_password",
	}

	// Настройка поведения моков
	userRepo.On("Exists", ctx, userEmail).Return(false, nil) // пользователь не существует
	userRepo.On("Create", ctx, newUser).Return(nil)          // успешная регистрация пользователя

	// Создание UseCase
	useCase := NewUserUseCase(userRepo, logger)

	// Вызов метода
	err := useCase.RegisterUser(ctx, newUser)

	// Проверки
	assert.NoError(t, err)
	userRepo.AssertCalled(t, "Create", ctx, newUser)
}

func TestRegisterUser_Error_UserAlreadyExists(t *testing.T) {
	ctx := context.Background()
	userEmail := "test@example.com"

	// Моки
	userRepo := mockUser.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Пример пользователя
	newUser := models.User{
		ID:    uuid.New(),
		Email: userEmail,
		Role:  "user",
		Hash:  "hashed_password",
	}

	// Настройка поведения моков
	userRepo.On("Exists", ctx, userEmail).Return(true, nil) // пользователь уже существует

	// Создание UseCase
	useCase := NewUserUseCase(userRepo, logger)

	// Вызов метода
	err := useCase.RegisterUser(ctx, newUser)

	// Проверки
	assert.Error(t, err)
	assert.Equal(t, "user already exists", err.Error())
	userRepo.AssertCalled(t, "Exists", ctx, userEmail)
	userRepo.AssertNotCalled(t, "Create", ctx, newUser)
}

func TestGetUserByEmail_Success(t *testing.T) {
	ctx := context.Background()
	userEmail := "test@example.com"

	// Моки
	userRepo := mockUser.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Пример пользователя
	expectedUser := models.User{
		ID:    uuid.New(),
		Email: userEmail,
		Role:  "user",
		Hash:  "hashed_password",
	}

	// Настройка поведения моков
	userRepo.On("GetByEmail", ctx, userEmail).Return(expectedUser, nil) // успешное получение пользователя

	// Создание UseCase
	useCase := NewUserUseCase(userRepo, logger)

	// Вызов метода
	user, err := useCase.GetUserByEmail(ctx, userEmail)

	// Проверки
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
	userRepo.AssertCalled(t, "GetByEmail", ctx, userEmail)
}

func TestGetUserByEmail_Error(t *testing.T) {
	ctx := context.Background()
	userEmail := "test@example.com"

	// Моки
	userRepo := mockUser.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Настройка поведения моков
	userRepo.On("GetByEmail", ctx, userEmail).Return(models.User{}, errors.New("user not found")) // ошибка при получении пользователя

	// Создание UseCase
	useCase := NewUserUseCase(userRepo, logger)

	// Вызов метода
	user, err := useCase.GetUserByEmail(ctx, userEmail)

	// Проверки
	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())
	assert.Equal(t, models.User{}, user)
	userRepo.AssertCalled(t, "GetByEmail", ctx, userEmail)
}

func TestCheckUserExists_Success(t *testing.T) {
	ctx := context.Background()
	userEmail := "test@example.com"

	// Моки
	userRepo := mockUser.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Настройка поведения моков
	userRepo.On("Exists", ctx, userEmail).Return(true, nil) // пользователь существует

	// Создание UseCase
	useCase := NewUserUseCase(userRepo, logger)

	// Вызов метода
	exists, err := useCase.CheckUserExists(ctx, userEmail)

	// Проверки
	assert.NoError(t, err)
	assert.True(t, exists)
	userRepo.AssertCalled(t, "Exists", ctx, userEmail)
}

func TestCheckUserExists_Error(t *testing.T) {
	ctx := context.Background()
	userEmail := "test@example.com"

	// Моки
	userRepo := mockUser.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Настройка поведения моков
	userRepo.On("Exists", ctx, userEmail).Return(false, errors.New("error checking user existence")) // ошибка при проверке существования

	// Создание UseCase
	useCase := NewUserUseCase(userRepo, logger)

	// Вызов метода
	exists, err := useCase.CheckUserExists(ctx, userEmail)

	// Проверки
	assert.Error(t, err)
	assert.False(t, exists)
	userRepo.AssertCalled(t, "Exists", ctx, userEmail)
}
