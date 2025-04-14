package service

import (
	"fmt"
	"github/ledger/account"
	"time"
)

const ACCOUNT_CAP = 100

type Service struct {
	accounts map[int64]*account.Account
}

func New() *Service {
	return &Service{
		accounts: make(map[int64]*account.Account, ACCOUNT_CAP),
	}
}

func (s *Service) CreateAccount(name string) int64 {
	acc := account.New(name)
	s.accounts[acc.ID] = acc
	return acc.ID
}

func (s *Service) RecordTransaction(accountID int64, date time.Time, transactionType string, amount float64) error {

	if _, ok := s.accounts[accountID]; !ok {
		return fmt.Errorf("account does not exist: %+v", accountID)
	}

	acc := s.accounts[accountID]

	err := acc.AddTransaction(transactionType, amount, date)
	if err != nil {
		return fmt.Errorf("error recording transaction: %+v", err)
	}

	return nil
}

func (s *Service) GetBalance(accountID int64) (float64, error) {

	if _, ok := s.accounts[accountID]; !ok {
		return 0, fmt.Errorf("account does not exist: %+v", accountID)
	}

	acc := s.accounts[accountID]

	return acc.Balance, nil
}

func (s *Service) GetTransactions(accountID int64) ([]*account.Transaction, error) {
	if _, ok := s.accounts[accountID]; !ok {
		return nil, fmt.Errorf("account does not exist: %+v", accountID)
	}

	acc := s.accounts[accountID]

	return acc.Transactions, nil
}
