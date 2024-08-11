package user

import (
	"errors"
	"testing"

	"github.com/personal/mini-aspire/internal/model"
	"github.com/personal/mini-aspire/internal/service/mappers"
	"github.com/petergtz/pegomock/v4"
)

func Test_AddUser(t *testing.T) {
	pegomock.RegisterMockTestingT(t)
	mockDatabase := model.NewMockIDatabase()
	service := NewService(mockDatabase)
	pegomock.When(mockDatabase.AddUser(&model.User{
		Name:     "test",
		Phone:    123,
		Password: "test",
	})).ThenReturn(nil)
	err := service.AddUser(nil, &mappers.UserRegistrationRequest{
		Name:     "test",
		Phone:    123,
		Password: "test",
	})

	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
}

func Test_AddUserInvalidRequest(t *testing.T) {
	pegomock.RegisterMockTestingT(t)
	mockDatabase := model.NewMockIDatabase()
	service := NewService(mockDatabase)
	err := service.AddUser(nil, &mappers.UserRegistrationRequest{
		Name:     "",
		Phone:    0,
		Password: "",
	})

	if err == nil {
		t.Errorf("expected error, got %v", err)
	}
}

func Test_GetUser(t *testing.T) {
	pegomock.RegisterMockTestingT(t)
	mockDatabase := model.NewMockIDatabase()
	service := NewService(mockDatabase)
	pegomock.When(mockDatabase.GetUser(map[string]interface{}{
		"id": uint64(1),
	})).ThenReturn(&model.User{
		ID:    1,
		Name:  "test",
		Phone: 123,
	}, nil)
	user, err := service.GetUser(nil, 1)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}

	if user == nil {
		t.Errorf("expected user, got nil")
	}
}

func Test_GetUserInvalidRequest(t *testing.T) {
	pegomock.RegisterMockTestingT(t)
	mockDatabase := model.NewMockIDatabase()
	service := NewService(mockDatabase)
	user, err := service.GetUser(nil, 0)
	if err == nil {
		t.Errorf("expected error, got %v", err)
	}

	if user != nil {
		t.Errorf("expected nil user, got %v", user)
	}
}

func Test_GetUserNotFound(t *testing.T) {
	pegomock.RegisterMockTestingT(t)
	mockDatabase := model.NewMockIDatabase()
	service := NewService(mockDatabase)
	pegomock.When(mockDatabase.GetUser(map[string]interface{}{
		"id": uint64(1),
	})).ThenReturn(nil, nil)
	user, err := service.GetUser(nil, 1)
	if err == nil {
		t.Errorf("expected error, got %v", err)
	}

	if user != nil {
		t.Errorf("expected nil user, got %v", user)
	}
}

func Test_GetUserDatabaseError(t *testing.T) {
	pegomock.RegisterMockTestingT(t)
	mockDatabase := model.NewMockIDatabase()
	service := NewService(mockDatabase)
	pegomock.When(mockDatabase.GetUser(map[string]interface{}{
		"id": uint64(1),
	})).ThenReturn(nil, errors.New("database error"))
	user, err := service.GetUser(nil, 1)
	if err == nil {
		t.Errorf("expected error, got %v", err)
	}

	if user != nil {
		t.Errorf("expected nil user, got %v", user)
	}
}
