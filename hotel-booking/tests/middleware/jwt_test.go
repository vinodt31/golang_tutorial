package middleware_test

import (
	"hotel-booking/internal/middleware"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

const testSecret = "my_super_secret_test_key"

// Helper function to generate a real signed token
func createTestToken(t *testing.T) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": 1,
		"email":   "vinod@gmail.com",
		"exp":     time.Now().Add(time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(testSecret))
	assert.NoError(t, err)
	return tokenString
}

func TestMissingToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.Use(middleware.JWTAuth())

	router.GET("/profile", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	req, _ := http.NewRequest("GET", "/profile", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestValidToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Inject the matching environment variable into the test runtime scope
	t.Setenv("JWT_SECRET", testSecret)

	router.Use(middleware.JWTAuth())

	router.GET("/profile", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	validToken := createTestToken(t)

	req, _ := http.NewRequest("GET", "/profile", nil)
	req.Header.Set("Authorization", "Bearer "+validToken) // "Bearer " prefix is parsed safely now

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "success")
}

func TestInvalidToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.Use(middleware.JWTAuth())

	router.GET("/profile", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	})

	req, _ := http.NewRequest("GET", "/profile", nil)
	req.Header.Set("Authorization", "Bearer invalid-token-string")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
