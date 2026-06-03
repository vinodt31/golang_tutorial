package service

import (
	"errors"
	"os"
	"time"

	"hotel-booking/internal/model"
	"hotel-booking/internal/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Repo repository.UserRepositoryInterface
}

func NewAuthService(
	repo repository.UserRepositoryInterface,
) *AuthService {

	return &AuthService{
		Repo: repo,
	}
}

func (s *AuthService) Register(
	name string,
	email string,
	password string,
) error {
	// 1. Check if user already exists
	existingUser, err := s.Repo.FindByEmail(email)
	if err == nil && existingUser != nil {
		return errors.New("email already registered")
	}

	// 2. Proceed if email is available
	hash, _ := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	user := model.User{
		Name:     name,
		Email:    email,
		Password: string(hash),
	}

	return s.Repo.Create(&user)
}

func (s *AuthService) Login(
	email string,
	password string,
) (string, error) {

	user, err := s.Repo.FindByEmail(email)

	if err != nil {
		return "", errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)

	if err != nil {
		return "", errors.New("invalid password")
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp": time.Now().
			Add(time.Hour * 24).
			Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(
		[]byte(os.Getenv("JWT_SECRET")),
	)
}
