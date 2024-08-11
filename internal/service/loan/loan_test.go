package loan

import (
	"testing"

	"github.com/personal/mini-aspire/internal/model"
	time2 "github.com/personal/mini-aspire/internal/pkg/time"
	"github.com/personal/mini-aspire/internal/service/admin"
	"github.com/personal/mini-aspire/internal/service/domains"
	"github.com/personal/mini-aspire/internal/service/mappers"
	"github.com/personal/mini-aspire/internal/service/user"
	"github.com/petergtz/pegomock/v4"
)

func Test_SubmitLoanRequest(t *testing.T) {
	pegomock.RegisterMockTestingT(t)
	mockDatabase := model.NewMockIDatabase()
	mockAdmin := admin.NewMockIAdmin()
	mockUser := user.NewMockIUser()
	service := NewService(mockDatabase, mockUser, mockAdmin)
	pegomock.When(mockDatabase.AddLoanRequest(&model.LoanApplication{
		UserID: 1,
		Amount: 1000,
		Terms:  1,
	})).ThenReturn(nil)
	err := service.SubmitLoanRequest(nil, &mappers.LoanRequest{
		UserID: 1,
		Amount: 1000,
		Term:   1,
		Date:   "2021-01-01",
	})

	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
}

func Test_SubmitLoanRequestInvalidRequest(t *testing.T) {
	pegomock.RegisterMockTestingT(t)
	mockDatabase := model.NewMockIDatabase()
	mockAdmin := admin.NewMockIAdmin()
	mockUser := user.NewMockIUser()
	service := NewService(mockDatabase, mockUser, mockAdmin)
	err := service.SubmitLoanRequest(nil, &mappers.LoanRequest{
		UserID: 1,
		Amount: 0,
		Term:   1,
		Date:   "2021-01-01",
	})

	if err == nil {
		t.Errorf("expected error, got %v", err)
	}
}

func Test_GetLoans(t *testing.T) {
	pegomock.RegisterMockTestingT(t)
	mockDatabase := model.NewMockIDatabase()
	mockAdmin := admin.NewMockIAdmin()
	mockUser := user.NewMockIUser()
	service := NewService(mockDatabase, mockUser, mockAdmin)
	pegomock.When(mockUser.GetUser(nil, uint64(1))).ThenReturn(&domains.User{
		ID: 1,
	}, nil)
	pegomock.When(mockDatabase.GetLoans(map[string]interface{}{
		"user_id": uint64(1),
	})).ThenReturn([]*model.LoanApplication{
		{
			UserID: 1,
			Amount: 1000,
			Terms:  1,
		},
	}, nil)
	loans, err := service.GetLoans(nil, 1)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}

	if loans == nil {
		t.Errorf("expected loans, got nil")
	}
}

func Test_GetLoansInvalidRequest(t *testing.T) {
	pegomock.RegisterMockTestingT(t)
	mockDatabase := model.NewMockIDatabase()
	mockAdmin := admin.NewMockIAdmin()
	mockUser := user.NewMockIUser()
	service := NewService(mockDatabase, mockUser, mockAdmin)
	_, err := service.GetLoans(nil, 0)
	if err == nil {
		t.Errorf("expected error, got %v", err)
	}
}

func Test_RepayLoanInstallment(t *testing.T) {
	pegomock.RegisterMockTestingT(t)
	mockDatabase := model.NewMockIDatabase()
	mockAdmin := admin.NewMockIAdmin()
	mockUser := user.NewMockIUser()
	timeService := time2.NewService()
	currentDate, _ := timeService.GetCurrentDate()
	service := NewService(mockDatabase, mockUser, mockAdmin)
	pegomock.When(mockDatabase.GetLoanDetails(uint64(1))).ThenReturn(&model.LoanApplication{
		ID:     1,
		Amount: 1000,
		Terms:  1,
		UserID: 1,
		Status: domains.LoanStatusApproved,
		LoanRepaymentData: []*model.LoanRepayment{
			{
				ID:                1,
				LoanApplicationID: 1,
				InstallmentAmount: 1000,
				PaidAmount:        0,
				DueDate:           currentDate,
				Status:            domains.LoanRepaymentStatusPending,
			},
		},
	}, nil)
	pegomock.When(mockDatabase.UpdateLoanRequest(&model.LoanApplication{
		ID: 1,
	})).ThenReturn(nil)
	pegomock.When(mockUser.GetUser(nil, uint64(1))).ThenReturn(&domains.User{
		ID: 1,
	}, nil)
	err := service.RepayLoanInstallment(nil, &mappers.RepaymentRequest{
		LoanApplicationID: 1,
		Amount:            1000,
		LoanRepaymentID:   1,
		UserID:            1,
	})

	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
}

func Test_RepayLoanInstallmentInvalidRequest(t *testing.T) {
	pegomock.RegisterMockTestingT(t)
	mockDatabase := model.NewMockIDatabase()
	mockAdmin := admin.NewMockIAdmin()
	mockUser := user.NewMockIUser()
	service := NewService(mockDatabase, mockUser, mockAdmin)
	err := service.RepayLoanInstallment(nil, &mappers.RepaymentRequest{
		LoanApplicationID: 1,
		Amount:            1000,
		LoanRepaymentID:   1,
		UserID:            1,
	})

	if err == nil {
		t.Errorf("expected error, got %v", err)
	}
}

// Test_calculateAmountEachTerm tests the calculateAmountEachTerm function
func Test_calculateAmountEachTerm(t *testing.T) {
	service := NewService(nil, nil, nil)
	type args struct {
		amount float64
		term   uint64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "Test case 1",
			args: args{
				amount: 1000,
				term:   1,
			},
			want: []float64{1000.0},
		},
		{
			name: "Test case 2",
			args: args{
				amount: 1000,
				term:   2,
			},
			want: []float64{500.0, 500.0},
		},
		{
			name: "Test case 3",
			args: args{
				amount: 1000,
				term:   3,
			},
			want: []float64{333.33, 333.33, 333.34},
		},
		{
			name: "Test case 4",
			args: args{
				amount: 0,
				term:   3,
			},
			want: []float64{},
		},
		{
			name: "Test case 5",
			args: args{
				amount: 1000,
				term:   0,
			},
			want: []float64{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := service.calculateAmountEachTerm(tt.args.amount, int(tt.args.term))
			for i := 0; i < len(got); i++ {
				if got[i] != tt.want[i] {
					t.Errorf("calculateAmountEachTerm() = %v, want %v", got[i], tt.want[i])
					t.Fail()
				}
			}
		})
	}
}
