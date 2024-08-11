package loan

import (
	"context"

	"github.com/personal/mini-aspire/internal/model"
	"github.com/personal/mini-aspire/internal/service/admin"
	"github.com/personal/mini-aspire/internal/service/mappers"
	"github.com/personal/mini-aspire/internal/service/user"
)

type ILoan interface {
	SubmitLoanRequest(context.Context, *mappers.LoanRequest) error
	GetLoans(context.Context, uint64) ([]*mappers.LoanRequest, error)
	RepayLoanInstallment(context.Context, *mappers.RepaymentRequest) error
}

type Service struct {
	database     model.IDatabase
	userService  user.IUser
	adminService admin.IAdmin
}

func NewService(
	database model.IDatabase,
	userService user.IUser,
	adminService admin.IAdmin,
) *Service {
	return &Service{
		database:     database,
		userService:  userService,
		adminService: adminService,
	}
}
