package admin

import (
	"context"
	"errors"
	"log"

	"github.com/personal/mini-aspire/internal/model"
	"github.com/personal/mini-aspire/internal/service/domains"
	"github.com/personal/mini-aspire/internal/service/mappers"
)

// AddAdmin adds a new admin to the system.
func (s *Service) AddAdmin(ctx context.Context, request *mappers.AdminRegistrationRequest) error {
	if valid := request.IsValid(); !valid {
		log.Printf("invalid request: %v", request)
		return errors.New("invalid request")
	}

	err := s.database.AddAdmin(&model.Admin{
		Name:     request.Name,
		Password: request.Password,
	})
	if err != nil {
		return err
	}

	return nil
}

// GetAdmin returns the admin details based on the admin ID.
func (s *Service) GetAdmin(ctx context.Context, adminID uint64) (*domains.Admin, error) {
	if adminID == 0 {
		return nil, errors.New("invalid admin ID")
	}

	admin, err := s.database.GetAdmin(map[string]interface{}{
		"id": adminID,
	})
	if err != nil {
		log.Printf("error getting admin: %v", err)
		return nil, err
	}

	if admin == nil || admin.ID == 0 {
		return nil, errors.New("admin not found")
	}

	return &domains.Admin{
		ID:   admin.ID,
		Name: admin.Name,
	}, nil
}

// GetSubmittedLoanRequests returns all the loan requests submitted by the users. For simplicity, I have considered
// only the submitted loan requests. In a real-world scenario, we can have more filters like pending, approved, rejected, etc.
func (s *Service) GetSubmittedLoanRequests(ctx context.Context, request *mappers.AdminLoanRequest) ([]*mappers.LoanRequest, error) {
	if request == nil || !request.IsValid() {
		return nil, errors.New("invalid request")
	}

	// check if the admin is valid
	admin, err := s.database.GetAdmin(map[string]interface{}{
		"id": request.AdminID,
	})
	if err != nil {
		return nil, err
	}

	if admin == nil || admin.ID == 0 {
		return nil, errors.New("admin not found")
	}

	// get all the loan requests
	loanRequests, err := s.database.GetLoans(map[string]interface{}{
		"status": domains.LoanStatusSubmitted,
	})

	if err != nil {
		return nil, err
	}

	loanRequestMappers := make([]*mappers.LoanRequest, 0, len(loanRequests))
	for _, loanRequest := range loanRequests {
		loanRequestMappers = append(loanRequestMappers, &mappers.LoanRequest{
			LoanID: loanRequest.ID,
			UserID: loanRequest.UserID,
			Amount: loanRequest.Amount,
			Date:   loanRequest.Date.Format("2006-01-02"),
			Term:   loanRequest.Terms,
		})
	}

	return loanRequestMappers, nil
}

// UpdateLoan updates the loan status based on the admin's decision.
func (s *Service) UpdateLoan(ctx context.Context, request *mappers.AdminLoanUpdateRequest) error {
	if request == nil || !request.IsValid() {
		return errors.New("invalid request")
	}

	// get the loan details
	loan, err := s.database.GetLoanDetails(request.LoanID)
	if err != nil {
		return err
	}

	// validate the loan for update
	err = s.validateLoanForUpdate(ctx, loan, request)
	if err != nil {
		return err
	}

	// update the loan status
	loan.Status = domains.LoanStatus(request.Status)
	err = s.database.UpdateLoanRequest(loan)
	if err != nil {
		return err
	}

	// if a loan is being cancelled or rejected, we need to change the state of the loan repayments as well
	if loan.Status == domains.LoanStatusCancelled || loan.Status == domains.LoanStatusRejected {
		for _, repayment := range loan.LoanRepaymentData {
			repayment.Status = domains.LoanRepaymentStatusCancelled
			err = s.database.UpdateLoanRepayment(
				map[string]interface{}{
					"id": repayment.ID,
				},
				repayment,
			)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// validateLoanForUpdate validates the loan for update. The loan should be in the submitted state and the loan ID should match.
// The admin should also be valid.
func (s *Service) validateLoanForUpdate(
	ctx context.Context,
	loan *model.LoanApplication,
	request *mappers.AdminLoanUpdateRequest,
) error {
	if loan == nil || loan.ID == 0 {
		return errors.New("loan not found")
	}

	if loan.Status != domains.LoanStatusSubmitted {
		return errors.New("loan is not in submitted state")
	}

	if loan.ID != request.LoanID {
		return errors.New("loan id mismatch")
	}

	// check if the admin is valid
	admin, err := s.database.GetAdmin(map[string]interface{}{
		"id": request.AdminID,
	})
	if err != nil {
		return err
	}

	if admin == nil || admin.ID == 0 {
		return errors.New("admin not found")
	}

	return nil
}
