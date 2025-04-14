package service

import (
	"github/ledger/account"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	s := New()

	assert.Equal(t, &Service{
		accounts: make(map[int64]*account.Account, ACCOUNT_CAP),
	}, s)
}

func TestCreateAccount(t *testing.T) {
	s := New()
	id := s.CreateAccount("AA")
	assert.Equal(t, s.accounts[id].Balance, 0)
}

func TestRecordTransaction(t *testing.T) {
	type input struct {
		accountID       int64
		transactionType string
		amount          float64
		date            time.Time
	}

	tests := []struct {
		name        string
		input       input
		expectedErr bool
	}{
		{
			name: "not enough funds for withdrawal",
			input: input{
				transactionType: account.TRANSACTION_WITHDRAWAL,
				amount:          500,
				date:            time.Now(),
			},
			expectedErr: true,
		},
		{
			name: "account id invalid",
			input: input{
				accountID:       10,
				transactionType: account.TRANSACTION_WITHDRAWAL,
				amount:          500,
				date:            time.Now(),
			},
			expectedErr: true,
		},
		{
			name: "transaction was successful",
			input: input{
				transactionType: account.TRANSACTION_DEPOSIT,
				amount:          500,
				date:            time.Now(),
			},
			expectedErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			name := "AA"
			s := New()
			id := s.CreateAccount(name)
			if tt.input.accountID != 0 {
				id = tt.input.accountID
			}
			err := s.RecordTransaction(id, tt.input.date, tt.input.transactionType, tt.input.amount)
			assert.Equal(t, tt.expectedErr, err != nil)
		})
	}
}

func TestGetBalance(t *testing.T) {

	tests := []struct {
		name            string
		accountID       int64
		expectedBalance float64
		expectedErr     bool
	}{
		{
			name:            "invalid account id",
			accountID:       10,
			expectedBalance: 0,
			expectedErr:     true,
		},
		{
			name:            "valid account id",
			expectedBalance: 0,
			expectedErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			name := "AA"
			s := New()
			id := s.CreateAccount(name)
			if tt.accountID != 0 {
				id = tt.accountID
			}
			balance, err := s.GetBalance(id)
			assert.Equal(t, tt.expectedErr, err != nil)
			assert.Equal(t, tt.expectedBalance, balance)
		})
	}
}

func TestGetTransactions(t *testing.T) {

	tests := []struct {
		name                 string
		accountID            int64
		expectedTransactions []*account.Transaction
		expectedErr          bool
	}{
		{
			name:        "invalid account id",
			accountID:   10,
			expectedErr: true,
		},
		{
			name:                 "valid account id",
			expectedTransactions: make([]*account.Transaction, 0, account.TRANSACTION_CAP),
			expectedErr:          false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			name := "AA"
			s := New()
			id := s.CreateAccount(name)
			if tt.accountID != 0 {
				id = tt.accountID
			}
			transactions, err := s.GetTransactions(id)
			assert.Equal(t, tt.expectedErr, err != nil)
			assert.Equal(t, tt.expectedTransactions, transactions)
		})
	}
}
