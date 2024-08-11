package loan

import (
	"context"
	"errors"
	"log"
	"math"
	"slices"
	"time"

	"github.com/personal/mini-aspire/internal/model"
	time2 "github.com/personal/mini-aspire/internal/pkg/time"
	"github.com/personal/mini-aspire/internal/service/domains"
	"github.com/personal/mini-aspire/internal/service/mappers"
)

const (
	DateFormat         = "2006-01-02"
	RepaymentFrequency = 7 // this is currently kept as constant, this can be accepted in user's input
)

func (s *Service) SubmitLoanRequest(ctx context.Context, request *mappers.LoanRequest) error {
	if valid := request.IsValid(); !valid {
		log.Printf("invalid request: %v", request)
		return errors.New("invalid request")
	}

	date, parseErr := time.Parse(DateFormat, request.Date)
	if parseErr != nil {
		log.Printf("could not parse Date from input: %v", parseErr)
		return errors.New("could not parse Date from input")
	}

	// Calculate the amount to be paid each term
	amountEachTerm := s.calculateAmountEachTerm(request.Amount, int(request.Term))
	loanRepaymentModel := make([]*model.LoanRepayment, 0, len(amountEachTerm))
	dateForPayment := date
	for i, amount := range amountEachTerm {
		dateForPayment = date.AddDate(0, 0, i*RepaymentFrequency)
		loanRepaymentModel = append(loanRepaymentModel, &model.LoanRepayment{
			InstallmentAmount: amount,
			PaidAmount:        0,
			DueDate:           dateForPayment,
			//PaidDate:          time.Time{},
			Status: domains.LoanRepaymentStatusPending,
		})
	}
	// Add loan request to database
	err := s.database.AddLoanRequest(&model.LoanApplication{
		UserID:            request.UserID,
		Amount:            request.Amount,
		Date:              date,
		Terms:             request.Term,
		Status:            domains.LoanStatusSubmitted,
		LoanRepaymentData: loanRepaymentModel,
	})

	if err != nil {
		log.Printf("error adding loan request: %v", err)
		return err
	}

	return nil
}

func (s *Service) GetLoans(ctx context.Context, userID uint64) ([]*mappers.LoanRequest, error) {
	if userID == 0 {
		return nil, errors.New("invalid user ID")
	}

	// validate if user exists in our system
	user, err := s.userService.GetUser(ctx, userID)
	if err != nil {
		log.Printf("error getting user: %v", err)
		return nil, err
	}

	if user == nil || user.ID == 0 {
		return nil, errors.New("user not found")
	}

	loans, err := s.database.GetLoans(map[string]interface{}{
		"user_id": userID,
	})
	if err != nil {
		log.Printf("error getting loans: %v", err)
		return nil, err
	}

	loanRequests := make([]*mappers.LoanRequest, 0, len(loans))
	for _, loan := range loans {
		loanRepayments := make([]*mappers.LoanRepayment, 0, len(loan.LoanRepaymentData))
		for _, loanRepayment := range loan.LoanRepaymentData {
			loanRepaymentRequest := &mappers.LoanRepayment{
				ID:                loanRepayment.ID,
				LoanApplicationID: loanRepayment.LoanApplicationID,
				InstallmentAmount: loanRepayment.InstallmentAmount,
				PaidAmount:        loanRepayment.PaidAmount,
				DueDate:           loanRepayment.DueDate.Format(DateFormat),
				PaidDate:          loanRepayment.PaidDate.Format(DateFormat),
				Status:            uint8(loanRepayment.Status),
			}
			loanRepayments = append(loanRepayments, loanRepaymentRequest)
		}

		loanRequest := &mappers.LoanRequest{
			LoanID:        loan.ID,
			UserID:        loan.UserID,
			Amount:        loan.Amount,
			Date:          loan.Date.Format(DateFormat),
			Term:          loan.Terms,
			Status:        uint8(loan.Status),
			LoanRepayment: loanRepayments,
		}
		loanRequests = append(loanRequests, loanRequest)
	}

	return loanRequests, nil
}

func (s *Service) RepayLoanInstallment(ctx context.Context, request *mappers.RepaymentRequest) error {
	if valid := request.IsValid(); !valid {
		log.Printf("invalid request: %v", request)
		return errors.New("invalid request")
	}

	// Get loan details
	loan, err := s.database.GetLoanDetails(request.LoanApplicationID)
	if err != nil {
		log.Printf("error getting loan details: %v", err)
		return err
	}

	timeService := time2.NewService()
	currentDate, currentDateError := timeService.GetCurrentDate()
	if currentDateError != nil {
		log.Printf("error getting current date: %v", currentDateError)
		return currentDateError
	}

	// validate loan for repayment
	err = s.isLoanValidForRepayment(loan, request, currentDate)
	if err != nil {
		log.Printf("error validating loan for repayment: %v", err)
		return err
	}

	err = s.database.UpdateLoanRepayment(map[string]interface{}{
		"id":                  request.LoanRepaymentID,
		"loan_application_id": request.LoanApplicationID,
		"status":              domains.LoanRepaymentStatusPending,
	}, &model.LoanRepayment{
		InstallmentAmount: 0,
		PaidAmount:        request.Amount,
		PaidDate:          currentDate,
		Status:            domains.LoanRepaymentStatusPaid,
	})
	if err != nil {
		log.Printf("error updating loan repayment details: %v", err)
		return err
	}

	// check if all installments are paid, mark loan as paid. We already have the loan details from the GetLoanDetails
	// query above
	allPaid := true
	for _, loanRepayment := range loan.LoanRepaymentData {
		if loanRepayment.ID == request.LoanRepaymentID {
			continue // skip the current installment, as it is just paid successfully.
		}

		if loanRepayment.Status != domains.LoanRepaymentStatusPaid {
			allPaid = false
			break
		}
	}

	if allPaid {
		loan.Status = domains.LoanStatusPaid
		err = s.database.UpdateLoanRequest(loan)
		if err != nil {
			log.Printf("error updating loan details: %v", err)
			return err
		}
	}

	return nil
}

// calculateAmountEachTerm calculates the amount to be paid each term for a given loan amount and terms
func (s *Service) calculateAmountEachTerm(amount float64, terms int) []float64 {
	// safe check for invalid input. Although this is already validated at request level, but extra check can help
	// avoid issues if some other developer decides to call this function directly
	if amount <= 0 || terms <= 0 {
		return []float64{}
	}

	// Calculate the base value for each term
	base := math.Floor(amount/float64(terms)*100) / 100

	// Calculate the remainder after dividing the amount into terms
	remainder := amount - (base * float64(terms))

	// Create a slice to store the results
	result := make([]float64, terms)

	// Assign the base value to all terms
	for i := 0; i < terms; i++ {
		result[i] = base
	}

	// Distribute the remainder across the terms
	for i := 0; i < int(math.Round(remainder*100)); i++ {
		result[i] += 0.01
	}

	slices.Reverse(result)
	return result
}

func (s *Service) isLoanValidForRepayment(
	loanDetails *model.LoanApplication,
	request *mappers.RepaymentRequest,
	currentDate time.Time,
) error {
	if loanDetails == nil {
		return errors.New("loan not found")
	}

	// check if loan belongs to the user
	if loanDetails.UserID != request.UserID {
		return errors.New("loan does not belong to the user")
	}

	// Check if the loan is in approved status
	if loanDetails.Status != domains.LoanStatusApproved {
		return errors.New("loan not approved")
	}

	for _, loanRepayment := range loanDetails.LoanRepaymentData {
		if loanRepayment.ID != request.LoanRepaymentID {
			continue
		}

		// Check if the loan repayment is already paid
		if loanRepayment.Status == domains.LoanRepaymentStatusPaid {
			return errors.New("installment already paid")
		}

		// Check if the amount to be paid is not less than the installment amount. It can be greater than the installment amount
		// which will be adjusted in future installments
		if loanRepayment.InstallmentAmount > request.Amount {
			return errors.New("amount to be paid is less than the installment amount")
		}

		// if user is paying after the due date, I am adding 10% extra amount to the installment amount
		if currentDate.After(loanRepayment.DueDate) {
			loanRepayment.InstallmentAmount += loanRepayment.InstallmentAmount * 0.1
			if loanRepayment.InstallmentAmount < request.Amount {
				return errors.New("amount to be paid is less than the installment amount. Amount is increased by 10% due to late payment")
			}
		}
	}

	return nil
}
