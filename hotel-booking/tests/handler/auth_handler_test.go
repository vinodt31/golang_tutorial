package handler_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"hotel-booking/internal/handler"
	"hotel-booking/internal/service"
	"hotel-booking/tests/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterEndpoint_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockRepo := new(mocks.UserRepositoryMock)

	// 1. First, the service will look up if the email is already registered
	mockRepo.On("FindByEmail", "vinod@gmail.com").Return(nil, errors.New("user not found"))
	// 2. Since it is not found, it proceeds to call Create
	mockRepo.On("Create", mock.Anything).Return(nil)

	authService := service.NewAuthService(mockRepo)
	authHandler := handler.NewAuthHandler(authService)

	router.POST("/register", authHandler.Register)

	body := []byte(`{
        "name":"Vinod",
        "email":"vinod@gmail.com",
        "password":"123456"
    }`)

	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "registered")
	mockRepo.AssertExpectations(t)
}

func TestRegisterEndpoint_ValidationError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// 1. Create empty initialization parameters.
	// No mock definitions (.On()) are needed because they will never be reached!
	mockRepo := new(mocks.UserRepositoryMock)
	authService := service.NewAuthService(mockRepo)
	authHandler := handler.NewAuthHandler(authService)

	router.POST("/register", authHandler.Register)

	// Missing email and password fields
	body := []byte(`{"name":"Vinod"}`)

	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// 2. Assertions
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Confirms that zero database methods were called during this transaction lifecycle
	mockRepo.AssertExpectations(t)
}
