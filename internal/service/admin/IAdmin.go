package admin

import (
	"context"

	"github.com/personal/mini-aspire/internal/model"
	"github.com/personal/mini-aspire/internal/service/domains"
	"github.com/personal/mini-aspire/internal/service/mappers"
)

type IAdmin interface {
	AddAdmin(context.Context, *mappers.AdminRegistrationRequest) error
	GetAdmin(context.Context, uint64) (*domains.Admin, error)
	GetSubmittedLoanRequests(context.Context, *mappers.AdminLoanRequest) ([]*mappers.LoanRequest, error)
	UpdateLoan(context.Context, *mappers.AdminLoanUpdateRequest) error
}

type Service struct {
	database model.IDatabase
}

func NewService(database model.IDatabase) *Service {
	return &Service{
		database: database,
	}
}
