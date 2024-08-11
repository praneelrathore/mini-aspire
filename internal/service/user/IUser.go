package user

import (
	"context"

	"github.com/personal/mini-aspire/internal/model"
	"github.com/personal/mini-aspire/internal/service/domains"
	"github.com/personal/mini-aspire/internal/service/mappers"
)

type IUser interface {
	AddUser(context.Context, *mappers.UserRegistrationRequest) error
	GetUser(context.Context, uint64) (*domains.User, error)
}

type Service struct {
	database model.IDatabase
}

func NewService(database model.IDatabase) *Service {
	return &Service{
		database: database,
	}
}
