package transaction

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/williamchang80/sea-apd/domain/user"
	"github.com/williamchang80/sea-apd/dto/request/transaction"
)

type Transaction struct {
	gorm.Model
	Status     string `json:"status"`
	BankNumber string `json:"bank_number"`
	BankName   string `json:"bank_name"`
	Amount     int    `json:"amount"`
	User       user.User
	UserId     string `gorm:"ForeignKey:id"`
}

type TransactionUsecase interface {
	CreateTransaction(transaction.TransactionRequest) error
	UpdateTransactionStatus(transaction.UpdateTransactionRequest) error
}

type TransactionController interface {
	CreateTransaction(echo.Context) error
	UpdateTransactionStatus(echo.Context) error
}

type TransactionRepository interface {
	CreateTransaction(Transaction) error
	GetTransactionById(string) (*Transaction, error)
	UpdateTransactionStatus(string, string) error
}

func NewTransaction(status string, bankNumber string, bankName string, amount int) Transaction {
	return Transaction{
		Status:     status,
		BankNumber: bankNumber,
		BankName:   bankName,
		Amount:     amount,
	}
}