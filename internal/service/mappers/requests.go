package mappers

type UserRegistrationRequest struct {
	Name     string `json:"name,omitempty"`
	Phone    uint64 `json:"phone,omitempty"`
	Password string `json:"password,omitempty"`
}

type AdminRegistrationRequest struct {
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
}

type AdminLoanRequest struct {
	AdminID uint64 `json:"admin_id,omitempty"`
}

type AdminLoanUpdateRequest struct {
	LoanID  uint64 `json:"loan_id,omitempty"`
	AdminID uint64 `json:"admin_id,omitempty"`
	Status  uint8  `json:"status,omitempty"`
}

type LoanRequest struct {
	LoanID uint64  `json:"loan_id,omitempty"`
	UserID uint64  `json:"user_id,omitempty"`
	Amount float64 `json:"amount,omitempty"`
	Date   string  `json:"date,omitempty"`
	Term   uint64  `json:"term,omitempty"`
	Status uint8   `json:"status,omitempty"`
}

type RepaymentRequest struct {
	UserID            uint64  `json:"user_id,omitempty"`
	LoanApplicationID uint64  `json:"loan_application_id,omitempty"`
	Amount            float64 `json:"amount,omitempty"`
	LoanRepaymentID   uint64  `json:"loan_repayment_id,omitempty"`
}

func (u *UserRegistrationRequest) IsValid() bool {
	if u == nil {
		return false
	}

	if u.Name == "" || u.Phone == 0 || u.Password == "" {
		return false
	}

	return true
}

func (u *AdminRegistrationRequest) IsValid() bool {
	if u == nil {
		return false
	}

	if u.Name == "" || u.Password == "" {
		return false
	}

	return true
}

func (l *LoanRequest) IsValid() bool {
	if l == nil {
		return false
	}

	if l.UserID == 0 || l.Amount <= 0 || l.Date == "" || l.Term == 0 {
		return false
	}

	return true
}

func (r *RepaymentRequest) IsValid() bool {
	if r == nil {
		return false
	}

	if r.LoanApplicationID == 0 || r.Amount <= 0 || r.LoanRepaymentID == 0 {
		return false
	}

	return true
}

func (a *AdminLoanUpdateRequest) IsValid() bool {
	if a == nil {
		return false
	}

	if a.LoanID == 0 || a.AdminID == 0 || a.Status == 0 {
		return false
	}

	return true
}

func (a *AdminLoanRequest) IsValid() bool {
	if a == nil {
		return false
	}

	if a.AdminID == 0 {
		return false
	}

	return true
}
