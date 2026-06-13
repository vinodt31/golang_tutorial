package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"my-go-api/internal/middleware"
	"my-go-api/internal/models"
	"my-go-api/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	DB        *sql.DB
	UserRepo  repository.UserRepository
	JWTSecret []byte
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Platform string `json:"platform"` // Optional: "web", "mobile" etc. Defaults to "web"
}

type RegisterInput struct {
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=6"`
	PhoneNumber string `json:"phone_number"`
}

// Struct for explicit Logout Configuration payloads
type LogoutInput struct {
	Platform   string `json:"platform"`
	AllDevices bool   `json:"all_devices"`
}

func NewAuthHandler(db *sql.DB, userRepo repository.UserRepository, jwtSecret []byte) *AuthHandler {
	return &AuthHandler{
		DB:        db,
		UserRepo:  userRepo,
		JWTSecret: jwtSecret,
	}
}

// 1. REGISTER METHOD
func (h *AuthHandler) Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.UserRepo.GetByEmail(input.Email)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "A user with this email address already exists"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to secure user password"})
		return
	}

	newUser := &models.User{
		FirstName:    input.FirstName,
		LastName:     input.LastName,
		Email:        input.Email,
		PasswordHash: string(hashedPassword),
		PhoneNumber:  input.PhoneNumber,
	}

	if err := h.UserRepo.Create(newUser, input.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user to database"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user_id": newUser.ID,
	})
}

// 2. LOGIN METHOD (Saves active token sessions to DB)
func (h *AuthHandler) Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.UserRepo.GetByEmail(input.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if !user.IsActive {
		c.JSON(http.StatusForbidden, gin.H{"error": "This account is deactivated"})
		return
	}

	if input.Platform == "" {
		input.Platform = "web"
	}

	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &middleware.CustomClaims{
		UserID: user.ID,
		Roles:  user.Roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(h.JWTSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate authentication token"})
		return
	}

	// Persist the active session token down to database tracking table
	err = h.UserRepo.SaveSession(user.ID, tokenString, input.Platform, expirationTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initialize user session state"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

// 3. REFRESH TOKEN METHOD
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User context not found"})
		return
	}

	user, err := h.UserRepo.GetProfileByID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User profile not found"})
		return
	}

	if !user.IsActive {
		c.JSON(http.StatusForbidden, gin.H{"error": "This account has been deactivated"})
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &middleware.CustomClaims{
		UserID: user.ID,
		Roles:  user.Roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(h.JWTSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not refresh authentication token"})
		return
	}

	// Persist the refreshed token session string
	_ = h.UserRepo.SaveSession(user.ID, tokenString, "web", expirationTime)

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

// 4. LOGOUT METHOD (Removes token records completely)
func (h *AuthHandler) Logout(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	// Extract the actual token string from Authorization header to clean it out
	authHeader := c.GetHeader("Authorization")
	var currentToken string
	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		currentToken = authHeader[7:]
	}

	var input LogoutInput
	// Optional JSON parameters mapping binding check
	if err := c.ShouldBindJSON(&input); err != nil {
		input.Platform = "web"
		input.AllDevices = false
	}

	// Remove session records matching parameters from PostgreSQL database
	err := h.UserRepo.DeleteSession(userID, input.Platform, input.AllDevices, currentToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to terminate session status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged out successfully context synchronized",
	})
}

// 5. GET USER PROFILE METHOD
func (h *AuthHandler) GetUserProfile(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	userProfile, err := h.UserRepo.GetProfileByID(userID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "User profile not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load profile data"})
		return
	}

	c.JSON(http.StatusOK, userProfile)
}
