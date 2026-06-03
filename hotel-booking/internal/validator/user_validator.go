package validator

import (
	"errors"
	"hotel-booking/internal/dto"
	"strings"
)

func ValidateRegister(req dto.RegisterRequest) error {
	if strings.TrimSpace(req.Email) == "" {
		return errors.New("email is required")
	}

	if !strings.Contains(req.Email, "@") {
		return errors.New("invalid email")
	}

	if len(req.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}

	if strings.TrimSpace(req.Name) == "" {
		return errors.New("name is required")
	}

	return nil
}
