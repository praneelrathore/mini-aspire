package model

import (
	"log"
)

func (db *Database) AddUser(userModel *User) error {
	result := db.connection.DB.Model(&User{}).Create(userModel)

	if resultErr := result.Error; resultErr != nil {
		log.Printf("Could not create user: %+v, err: %v",
			userModel, resultErr)
		return resultErr
	}

	return nil
}

func (db *Database) GetUser(condition map[string]interface{}) (*User, error) {
	var userModel *User
	result := db.connection.DB.Where(condition).Find(&userModel)
	if resultErr := result.Error; resultErr != nil {
		log.Printf("Could not get user for condition %+v, err: %v", condition, resultErr)
		return nil, resultErr
	}

	return userModel, nil
}

func (db *Database) AddAdmin(adminModel *Admin) error {
	result := db.connection.DB.Model(&Admin{}).Create(adminModel)

	if resultErr := result.Error; resultErr != nil {
		log.Printf("Could not create admin: %+v, err: %v",
			adminModel, resultErr)
		return resultErr
	}

	return nil
}

func (db *Database) GetAdmin(condition map[string]interface{}) (*Admin, error) {
	var adminModel *Admin
	result := db.connection.DB.Where(condition).Find(&adminModel)

	if resultErr := result.Error; resultErr != nil {
		log.Printf("Could not get admin for condition %+v, err: %v", condition, resultErr)
		return nil, resultErr
	}

	return adminModel, nil
}

func (db *Database) AddLoanRequest(loanApplication *LoanApplication) error {
	result := db.connection.DB.Model(&LoanApplication{}).Create(loanApplication)

	if resultErr := result.Error; resultErr != nil {
		log.Printf("Could not create loan request: %+v, err: %v",
			loanApplication, resultErr)
		return resultErr
	}

	return nil
}

func (db *Database) UpdateLoanRequest(loanApplication *LoanApplication) error {
	result := db.connection.DB.Model(&LoanApplication{}).Where("id = ?", loanApplication.ID).Updates(loanApplication)

	if resultErr := result.Error; resultErr != nil {
		log.Printf("Could not update loan request: %+v, err: %v",
			loanApplication, resultErr)
		return resultErr
	}

	return nil
}

func (db *Database) GetLoans(condition map[string]interface{}) ([]*LoanApplication, error) {
	var loanApplications []*LoanApplication
	result := db.connection.DB.Preload("LoanRepaymentData").Where(condition).Find(&loanApplications)
	if resultErr := result.Error; resultErr != nil {
		log.Printf("Could not get loans for condition %+v, err: %v", condition, resultErr)
		return nil, resultErr
	}

	return loanApplications, nil
}

func (db *Database) GetLoanDetails(loanID uint64) (*LoanApplication, error) {
	var loanApplication *LoanApplication
	result := db.connection.DB.Preload("LoanRepaymentData").Where("id = ?", loanID).Find(&loanApplication)

	if resultErr := result.Error; resultErr != nil {
		log.Printf("Could not get loan details for loan ID %d, err: %v", loanID, resultErr)
		return nil, resultErr
	}

	return loanApplication, nil
}

func (db *Database) UpdateLoanRepayment(condition map[string]interface{}, loanRepayment *LoanRepayment) error {
	result := db.connection.DB.Model(&LoanRepayment{}).Where(condition).Updates(loanRepayment)

	if resultErr := result.Error; resultErr != nil {
		log.Printf("Could not update loan repayment for condition %+v, err: %v",
			condition, resultErr)
		return resultErr
	}

	return nil
}
