package domain

import (
	"banking/dto"
	"banking/errs"
)

const dbTSLayout = "2006-01-02 15:04:05"

type Account struct {
	AccountId   string  `json:"accountId" db:"account_id"`
	CustomerId  string  `json:"customerId" db:"customer_id"`
	OpeningDate string  `json:"openingDate" db:"opening_date"`
	AccountType string  `json:"accountType" db:"account_type"`
	Amount      float64 `json:"amount" db:"amount"`
	Status      string  `json:"status" db:"status"`
}

func (a Account) ToNewAccountResponseDto() *dto.NewAccountResponse {
	return &dto.NewAccountResponse{a.AccountId}
}

type AccountRepository interface {
	Save(account Account) (*Account, *errs.AppError)
	SaveTransaction(transaction Transaction) (*Transaction, *errs.AppError)
	FindBy(accountId string) (*Account, *errs.AppError)
}

func (a Account) CanWithdraw(amount float64) bool {
	if a.Amount < amount {
		return false
	}
	return true
}

func NewAccount(customerId, accountType string, amount float64) Account {
	return Account{
		CustomerId:  customerId,
		OpeningDate: dbTSLayout,
		AccountType: accountType,
		Amount:      amount,
		Status:      "1",
	}
}
