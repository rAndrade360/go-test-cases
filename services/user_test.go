package services

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/rAndrade360/go-mock-tests/repositories/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserReturningError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	name := "Renan"
	email := "renandotcorrea@gmail.com"

	mockUserRepository := mocks.NewMockUserRepository(mockCtrl)

	t.Run("Should not be able to create a user returning an error", func(t *testing.T) {
		mockUserRepository.EXPECT().Create(name, email).Return(errors.New("create user error"))

		userService := NewUserService(mockUserRepository)
		err := userService.Create(name, email)
		assert.EqualError(t, err, "create user error")
	})

	t.Run("Should be able to create a user", func(t *testing.T) {
		mockUserRepository.EXPECT().Create(name, email).Return(nil)

		userService := NewUserService(mockUserRepository)
		err := userService.Create(name, email)
		assert.Nil(t, err)
	})
}
