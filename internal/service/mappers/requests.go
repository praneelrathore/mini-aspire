package mappers

// UserRegistrationRequest is a struct that represents a user registration request
type UserRegistrationRequest struct {
	Name     string `json:"name,omitempty"`
	Phone    uint64 `json:"phone,omitempty"`
	Password string `json:"password,omitempty"`
}

// AdminRegistrationRequest is a struct that represents an admin registration request
type AdminRegistrationRequest struct {
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
}

// AdminLoanRequest is a struct that represents an admin loan request
type AdminLoanRequest struct {
	AdminID uint64 `json:"admin_id,omitempty"`
}

// AdminLoanUpdateRequest is a struct that represents an admin loan update request
type AdminLoanUpdateRequest struct {
	LoanID  uint64 `json:"loan_id,omitempty"`
	AdminID uint64 `json:"admin_id,omitempty"`
	Status  uint8  `json:"status,omitempty"`
}

// LoanRequest is a struct that represents a loan request
type LoanRequest struct {
	LoanID        uint64           `json:"loan_id,omitempty"`
	UserID        uint64           `json:"user_id,omitempty"`
	Amount        float64          `json:"amount,omitempty"`
	Date          string           `json:"date,omitempty"`
	Term          uint64           `json:"term,omitempty"`
	Status        uint8            `json:"status,omitempty"`
	LoanRepayment []*LoanRepayment `json:"loan_repayment,omitempty"`
}

// RepaymentRequest is a struct that represents a loan repayment request
type RepaymentRequest struct {
	UserID            uint64  `json:"user_id,omitempty"`
	LoanApplicationID uint64  `json:"loan_application_id,omitempty"`
	Amount            float64 `json:"amount,omitempty"`
	LoanRepaymentID   uint64  `json:"loan_repayment_id,omitempty"`
}

// LoanRepayment is a struct that represents a loan repayment request
type LoanRepayment struct {
	ID                uint64  `json:"id,omitempty"`
	LoanApplicationID uint64  `json:"loan_application_id,omitempty"`
	InstallmentAmount float64 `json:"installment_amount,omitempty"`
	PaidAmount        float64 `json:"paid_amount,omitempty"`
	DueDate           string  `json:"due_date,omitempty"`
	PaidDate          string  `json:"paid_date,omitempty"`
	Status            uint8   `json:"status,omitempty"`
}

// IsValid checks if the UserRegistrationRequest is valid
func (u *UserRegistrationRequest) IsValid() bool {
	if u == nil {
		return false
	}

	if u.Name == "" || u.Phone == 0 || u.Password == "" {
		return false
	}

	return true
}

// IsValid checks if the AdminRegistrationRequest is valid
func (u *AdminRegistrationRequest) IsValid() bool {
	if u == nil {
		return false
	}

	if u.Name == "" || u.Password == "" {
		return false
	}

	return true
}

// IsValid checks if the LoanRequest is valid
func (l *LoanRequest) IsValid() bool {
	if l == nil {
		return false
	}

	if l.UserID == 0 || l.Amount <= 0 || l.Date == "" || l.Term == 0 {
		return false
	}

	return true
}

// IsValid checks if the RepaymentRequest is valid
func (r *RepaymentRequest) IsValid() bool {
	if r == nil {
		return false
	}

	if r.LoanApplicationID == 0 || r.Amount <= 0 || r.LoanRepaymentID == 0 {
		return false
	}

	return true
}

// IsValid checks if the AdminLoanUpdateRequest is valid
func (a *AdminLoanUpdateRequest) IsValid() bool {
	if a == nil {
		return false
	}

	if a.LoanID == 0 || a.AdminID == 0 || a.Status == 0 {
		return false
	}

	return true
}

// IsValid checks if the AdminLoanRequest is valid
func (a *AdminLoanRequest) IsValid() bool {
	if a == nil {
		return false
	}

	if a.AdminID == 0 {
		return false
	}

	return true
}
