package mocks

import (
	"hotel-booking/internal/model"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) Create(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *UserRepositoryMock) FindByEmail(email string) (*model.User, error) {
	args := m.Called(email)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*model.User), args.Error(1)
}
