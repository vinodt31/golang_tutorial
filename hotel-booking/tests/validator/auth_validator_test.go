package validator_test

import (
	"hotel-booking/internal/dto"
	"hotel-booking/internal/validator"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateRegister(t *testing.T) {
	// Define the structural layout of our test table
	tests := []struct {
		name        string              // Name of the sub-test scenario
		request     dto.RegisterRequest // The input payload
		expectError bool                // Should this test return an error?
	}{
		{
			name: "Success case",
			request: dto.RegisterRequest{
				Name:     "Vinod",
				Email:    "vinod@gmail.com",
				Password: "123456",
			},
			expectError: false,
		},
		{
			name: "Invalid email syntax",
			request: dto.RegisterRequest{
				Name:     "Vinod",
				Email:    "vinodgmail.com",
				Password: "123456",
			},
			expectError: true,
		},
		{
			name: "Empty name",
			request: dto.RegisterRequest{
				Name:     "",
				Email:    "vinod@gmail.com",
				Password: "123456",
			},
			expectError: true,
		},
		{
			name: "Password too short",
			request: dto.RegisterRequest{
				Name:     "Vinod",
				Email:    "vinod@gmail.com",
				Password: "123",
			},
			expectError: true,
		},
	}

	// Iterate through each test case in the table
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validator.ValidateRegister(tc.request)

			if tc.expectError {
				assert.NotNil(t, err, "Expected an error but got nil")
			} else {
				assert.Nil(t, err, "Expected no error but got: %v", err)
			}
		})
	}
}
