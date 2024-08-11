package user

import (
	"context"
	"errors"
	"log"

	"github.com/personal/mini-aspire/internal/model"
	"github.com/personal/mini-aspire/internal/service/domains"
	"github.com/personal/mini-aspire/internal/service/mappers"
)

func (s *Service) AddUser(ctx context.Context, request *mappers.UserRegistrationRequest) error {
	if valid := request.IsValid(); !valid {
		log.Printf("invalid request: %v", request)
		return errors.New("invalid request")
	}

	err := s.database.AddUser(&model.User{
		Name:     request.Name,
		Phone:    request.Phone,
		Password: request.Password,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetUser(ctx context.Context, userID uint64) (*domains.User, error) {
	if userID == 0 {
		return nil, errors.New("invalid user ID")
	}

	user, err := s.database.GetUser(map[string]interface{}{
		"id": userID,
	})
	if err != nil {
		log.Printf("error getting user: %v", err)
		return nil, err
	}

	if user == nil || user.ID == 0 {
		return nil, errors.New("user not found")
	}

	return &domains.User{
		ID:    user.ID,
		Name:  user.Name,
		Phone: user.Phone,
	}, nil
}
