package model

import (
	"time"

	"github.com/personal/mini-aspire/internal/service/domains"
)

type User struct {
	ID        uint64 `gorm:"primary_key"`
	Name      string `gorm:"type:not null"`
	Phone     uint64 `gorm:"type:not null"`
	Password  string `gorm:"type:not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Admin struct {
	ID        uint64 `gorm:"primary_key"`
	Name      string `gorm:"type:not null"`
	Password  string `gorm:"type:not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type LoanApplication struct {
	ID                uint64  `gorm:"primary_key"`
	UserID            uint64  `gorm:"type:not null"`
	Amount            float64 `gorm:"type:not null"`
	Terms             uint64  `gorm:"type:not null"`
	Date              time.Time
	Status            domains.LoanStatus `gorm:"not null"`
	AdminID           uint64             `gorm:"default:null"`
	LoanRepaymentData []*LoanRepayment
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type LoanRepayment struct {
	ID                uint64  `gorm:"primary_key"`
	LoanApplicationID uint64  `gorm:"foreign_key:LoanApplication; not null"`
	InstallmentAmount float64 `gorm:"type:not null"`
	PaidAmount        float64 `gorm:"type:not null"`
	DueDate           time.Time
	PaidDate          time.Time                   `gorm:"default:null"`
	Status            domains.LoanRepaymentStatus `gorm:"not null"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
