package repository_test

import (
	"os"
	"testing"

	"hotel-booking/internal/config"
	"hotel-booking/internal/model"
	"hotel-booking/internal/repository"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "root") // Your local password
	os.Setenv("DB_NAME", "hotel_booking")
	os.Setenv("DB_SSLMODE", "disable")

	config.ConnectDB()

	os.Exit(m.Run())
}

func TestCreateUser(t *testing.T) {
	// 1. Start a temporary transaction
	tx := config.DB.Begin()

	// 2. Schedule an automatic Rollback when this test finishes
	defer tx.Rollback()

	// Pass the transaction 'tx' instead of 'config.DB'
	repo := repository.NewUserRepository(tx)

	user := model.User{
		Name:     "Test User",
		Email:    "transaction_test@gmail.com",
		Password: "123456",
	}

	err := repo.Create(&user)
	assert.Nil(t, err)

	// Verify it exists INSIDE the transaction context
	fetchedUser, err := repo.FindByEmail("transaction_test@gmail.com")
	assert.Nil(t, err)
	assert.NotNil(t, fetchedUser)

	// When the function ends, defer tx.Rollback() fires!
	// The data vanishes from your database completely.
}

func TestFindByEmail(t *testing.T) {
	tx := config.DB.Begin()
	defer tx.Rollback()

	repo := repository.NewUserRepository(tx)

	// Seed the user inside this transaction isolation block
	user := model.User{
		Name:     "Finder User",
		Email:    "finder_test@gmail.com",
		Password: "password123",
	}
	err := tx.Create(&user).Error
	assert.NoError(t, err)

	fetchedUser, err := repo.FindByEmail("finder_test@gmail.com")

	assert.Nil(t, err)
	assert.NotNil(t, fetchedUser)
	assert.Equal(t, "finder_test@gmail.com", fetchedUser.Email)
}
