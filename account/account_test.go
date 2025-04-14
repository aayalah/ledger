package account

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {

	name := "AA"

	acc := New(name)

	expectedAcc := &Account{
		ID:           0,
		Name:         name,
		Transactions: make([]*Transaction, 0, TRANSACTION_CAP),
	}

	assert.Equal(t, expectedAcc, acc)

}

func TestAddTransaction(t *testing.T) {

	type input struct {
		transactionType string
		amount          float64
		date            time.Time
	}

	tests := []struct {
		name           string
		input          input
		initialBalance float64
		expectedErr    bool
	}{
		{
			name: "not enough funds for withdrawal",
			input: input{
				transactionType: TRANSACTION_WITHDRAWAL,
				amount:          500,
				date:            time.Now(),
			},
			initialBalance: 10,
			expectedErr:    true,
		},
		{
			name: "valid deposit",
			input: input{
				transactionType: TRANSACTION_DEPOSIT,
				amount:          500,
				date:            time.Now(),
			},
			initialBalance: 0,
			expectedErr:    false,
		},
		{
			name: "valid withdrawal",
			input: input{
				transactionType: TRANSACTION_WITHDRAWAL,
				amount:          500,
				date:            time.Now(),
			},
			initialBalance: 1000,
			expectedErr:    false,
		},
		{
			name: "invalid transaction type",
			input: input{
				transactionType: "INVALID",
				amount:          500,
				date:            time.Now(),
			},
			initialBalance: 1000,
			expectedErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			name := "AA"
			acc := New(name)
			acc.Balance = tt.initialBalance
			err := acc.AddTransaction(tt.input.transactionType, tt.input.amount, tt.input.date)
			assert.Equal(t, tt.expectedErr, err != nil)
		})
	}
}
