package service_test

import (
	"errors"
	"hotel-booking/internal/model"
	"hotel-booking/internal/service"
	"hotel-booking/tests/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterSuccess(t *testing.T) {
	mockRepo := new(mocks.UserRepositoryMock)

	// Simulate that the email does NOT exist yet
	mockRepo.
		On("FindByEmail", "vinod@gmail.com").
		Return(nil, errors.New("user not found"))

	mockRepo.
		On("Create", mock.Anything).
		Return(nil)

	authService := service.NewAuthService(mockRepo)

	err := authService.Register("Vinod", "vinod@gmail.com", "123456")

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDuplicateEmail(t *testing.T) {

	mockRepo := new(mocks.UserRepositoryMock)

	mockRepo.
		On("FindByEmail", "vinod@gmail.com").
		Return(&model.User{
			Email: "vinod@gmail.com",
		}, nil)

	authService := service.NewAuthService(mockRepo)

	err := authService.Register(
		"Vinod",
		"vinod@gmail.com",
		"123456",
	)

	assert.NotNil(t, err)
}

func TestUserNotFound(t *testing.T) {

	mockRepo := new(mocks.UserRepositoryMock)

	mockRepo.
		On("FindByEmail", "abc@gmail.com").
		Return(nil, errors.New("not found"))

	authService := service.NewAuthService(mockRepo)

	_, err := authService.Login(
		"abc@gmail.com",
		"123456",
	)

	assert.NotNil(t, err)
}
