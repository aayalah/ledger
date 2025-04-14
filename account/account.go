package account

import (
	"fmt"
	"time"
)

type Transaction struct {
	Date            time.Time
	TransactionType string
	Amount          float64
}

type Account struct {
	ID           int64
	Name         string
	Balance      float64
	Transactions []*Transaction
}

const (
	TRANSACTION_WITHDRAWAL = "WITHDRAWAL"
	TRANSACTION_DEPOSIT    = "DEPOSIT"
)

var lastID int64 = 0

const TRANSACTION_CAP = 100

func New(name string) *Account {

	acc := &Account{
		ID:           lastID,
		Name:         name,
		Transactions: make([]*Transaction, 0, TRANSACTION_CAP),
	}

	lastID = lastID + 1

	return acc
}

func (acc *Account) AddTransaction(transactionType string, amount float64, date time.Time) error {

	if transactionType == TRANSACTION_WITHDRAWAL && amount > acc.Balance {
		return fmt.Errorf("not enough available funds for withdrawal: %+v", acc.ID)
	}

	switch transactionType {
	case TRANSACTION_DEPOSIT:
		acc.Balance = acc.Balance + amount
	case TRANSACTION_WITHDRAWAL:
		acc.Balance = acc.Balance - amount
	default:
		return fmt.Errorf("invalid transaction type: %+v", transactionType)
	}

	acc.Transactions = append(acc.Transactions, &Transaction{
		Date:            date,
		TransactionType: transactionType,
		Amount:          amount,
	})

	return nil
}
