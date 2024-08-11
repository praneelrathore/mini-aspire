package admin

import (
	"context"

	"github.com/personal/mini-aspire/internal/model"
	"github.com/personal/mini-aspire/internal/service/domains"
	"github.com/personal/mini-aspire/internal/service/mappers"
)

// IAdmin is an interface that defines the methods that must be implemented by an admin service.
type IAdmin interface {
	// AddAdmin adds a new admin to the database
	AddAdmin(context.Context, *mappers.AdminRegistrationRequest) error
	// GetAdmin retrieves an admin from the database
	GetAdmin(context.Context, uint64) (*domains.Admin, error)
	// GetSubmittedLoanRequests retrieves submitted loan requests from the database
	GetSubmittedLoanRequests(context.Context, *mappers.AdminLoanRequest) ([]*mappers.LoanRequest, error)
	// UpdateLoan updates an existing loan in the database
	UpdateLoan(context.Context, *mappers.AdminLoanUpdateRequest) error
}

// Service is a struct that represents an admin service object instance.
type Service struct {
	// database is a database model instance
	database model.IDatabase
}

// NewService creates a new admin service object instance
func NewService(database model.IDatabase) *Service {
	return &Service{
		database: database,
	}
}
