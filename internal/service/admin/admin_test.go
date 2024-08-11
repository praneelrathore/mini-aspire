package admin

import (
	"context"
	"testing"

	"github.com/personal/mini-aspire/internal/model"
	"github.com/personal/mini-aspire/internal/service/mappers"
	"github.com/petergtz/pegomock/v4"
)

func TestAdminService_AddAdmin(t *testing.T) {
	pegomock.RegisterMockTestingT(t)
	mockDatabase := model.NewMockIDatabase()
	adminService := NewService(mockDatabase)
	pegomock.When(mockDatabase.AddAdmin(&model.Admin{
		Name:     "test",
		Password: "test",
	})).ThenReturn(nil)
	err := adminService.AddAdmin(context.TODO(), &mappers.AdminRegistrationRequest{
		Name:     "test",
		Password: "test",
	})

	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
}

func TestAdminService_AddAdminInvalidRequest(t *testing.T) {
	pegomock.RegisterMockTestingT(t)
	mockDatabase := model.NewMockIDatabase()
	adminService := NewService(mockDatabase)
	err := adminService.AddAdmin(context.TODO(), &mappers.AdminRegistrationRequest{
		Name:     "",
		Password: "",
	})

	if err == nil {
		t.Errorf("expected error, got %v", err)
	}
}

func TestAdminService_GetAdmin(t *testing.T) {
	pegomock.RegisterMockTestingT(t)
	mockDatabase := model.NewMockIDatabase()
	adminService := NewService(mockDatabase)
	pegomock.When(mockDatabase.GetAdmin(map[string]interface{}{
		"id": uint64(1),
	})).ThenReturn(&model.Admin{
		ID:   1,
		Name: "test",
	}, nil)
	admin, err := adminService.GetAdmin(context.TODO(), 1)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}

	if admin == nil {
		t.Errorf("expected admin, got nil")
	}

	if admin.ID != 1 {
		t.Errorf("expected admin ID 1, got %d", admin.ID)
	}

	if admin.Name != "test" {
		t.Errorf("expected admin name test, got %s", admin.Name)
	}
}

func Test_GetSubmittedLoanRequests(t *testing.T) {
	pegomock.RegisterMockTestingT(t)
	mockDatabase := model.NewMockIDatabase()
	adminService := NewService(mockDatabase)
	pegomock.When(mockDatabase.GetAdmin(map[string]interface{}{
		"id": uint64(1),
	})).ThenReturn(&model.Admin{
		ID:   1,
		Name: "test",
	}, nil)
	pegomock.When(mockDatabase.GetLoans(map[string]interface{}{})).ThenReturn([]*model.LoanApplication{
		{
			ID:     1,
			UserID: 1,
			Amount: 1000,
			Terms:  1,
			Status: 1,
		},
	}, nil)
	loanRequests, err := adminService.GetSubmittedLoanRequests(context.TODO(), &mappers.AdminLoanRequest{
		AdminID: 1,
	})
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}

	if loanRequests == nil {
		t.Errorf("expected loan requests, got nil")
	}
}

func Test_GetSubmittedLoanRequestsInvalidRequest(t *testing.T) {
	pegomock.RegisterMockTestingT(t)
	mockDatabase := model.NewMockIDatabase()
	adminService := NewService(mockDatabase)
	loanRequests, err := adminService.GetSubmittedLoanRequests(context.TODO(), nil)
	if err == nil {
		t.Errorf("expected error, got %v", err)
	}

	if loanRequests != nil {
		t.Errorf("expected nil loan requests, got %v", loanRequests)
	}
}

func Test_GetSubmittedLoanRequestsInvalidAdmin(t *testing.T) {
	pegomock.RegisterMockTestingT(t)
	mockDatabase := model.NewMockIDatabase()
	adminService := NewService(mockDatabase)
	pegomock.When(mockDatabase.GetAdmin(map[string]interface{}{
		"id": uint64(1),
	})).ThenReturn(nil, nil)
	loanRequests, err := adminService.GetSubmittedLoanRequests(context.TODO(), &mappers.AdminLoanRequest{
		AdminID: 1,
	})
	if err == nil {
		t.Errorf("expected error, got %v", err)
	}

	if loanRequests != nil {
		t.Errorf("expected nil loan requests, got %v", loanRequests)
	}
}

func Test_UpdateLoan(t *testing.T) {
	pegomock.RegisterMockTestingT(t)
	mockDatabase := model.NewMockIDatabase()
	adminService := NewService(mockDatabase)
	pegomock.When(mockDatabase.GetLoanDetails(uint64(1))).ThenReturn(&model.LoanApplication{
		ID:     1,
		Status: 1,
	}, nil)
	pegomock.When(mockDatabase.GetAdmin(map[string]interface{}{
		"id": uint64(1),
	})).ThenReturn(&model.Admin{
		ID:   1,
		Name: "test",
	}, nil)
	pegomock.When(mockDatabase.UpdateLoanRequest(&model.LoanApplication{
		ID:     1,
		Status: 1,
	})).ThenReturn(nil)
	err := adminService.UpdateLoan(context.TODO(), &mappers.AdminLoanUpdateRequest{
		LoanID:  1,
		AdminID: 1,
		Status:  1,
	})
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
}
