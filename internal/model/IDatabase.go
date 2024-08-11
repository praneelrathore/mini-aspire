package model

import (
	"context"

	"github.com/personal/mini-aspire/internal/appInit"
)

// IDatabase is an interface that defines the methods that must be implemented by a database model.
type IDatabase interface {
	// AddUser adds a new user to the database
	AddUser(*User) error
	// GetUser retrieves a user from the database
	GetUser(map[string]interface{}) (*User, error)
	// AddAdmin adds a new admin to the database
	AddAdmin(*Admin) error
	// GetAdmin retrieves an admin from the database
	GetAdmin(map[string]interface{}) (*Admin, error)
	// AddLoanRequest adds a new loan request to the database
	AddLoanRequest(*LoanApplication) error
	// UpdateLoanRequest updates an existing loan request in the database
	UpdateLoanRequest(*LoanApplication) error
	// GetLoanDetails retrieves a loan request from the database
	GetLoanDetails(uint64) (*LoanApplication, error)
	// UpdateLoanRepayment updates a loan repayment to the database
	UpdateLoanRepayment(map[string]interface{}, *LoanRepayment) error
	// GetLoans retrieves loans from the database
	GetLoans(map[string]interface{}) ([]*LoanApplication, error)
}

// Database is a struct that represents a database object instance.
type Database struct {
	connection *appInit.Connection
}

// NewDatabase creates a new database object instance
func NewDatabase(ctx context.Context) IDatabase {
	appInit.InitializeConfig()
	db := appInit.InitializeDatabase(ctx)
	return &Database{connection: db}
}
