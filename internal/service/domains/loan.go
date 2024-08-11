package domains

import "time"

type LoanStatus uint8
type LoanRepaymentStatus uint8

const (
	LoanStatusSubmitted LoanStatus = 1
	LoanStatusApproved  LoanStatus = 2
	LoanStatusPaid      LoanStatus = 3
	LoanStatusRejected  LoanStatus = 4
	LoanStatusCancelled LoanStatus = 5

	LoanRepaymentStatusPending   LoanRepaymentStatus = 1
	LoanRepaymentStatusPaid      LoanRepaymentStatus = 2
	LoanRepaymentStatusSettled   LoanRepaymentStatus = 3
	LoanRepaymentStatusCancelled LoanRepaymentStatus = 4
)

type Loan struct {
	ID                uint64
	User              *User
	Admin             *Admin
	Amount            float64
	RepaymentSchedule []*RepaymentSchedule
	Status            LoanStatus
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type RepaymentSchedule struct {
	ID     uint64
	Date   time.Time
	Amount float64
}
