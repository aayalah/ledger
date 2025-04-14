package handlers

import (
	"bytes"
	"encoding/json"
	"github/ledger/account"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

type serviceMock struct {
}

func (sm *serviceMock) CreateAccount(name string) int64 {
	return 0
}

func (sm *serviceMock) RecordTransaction(accountID int64, date time.Time, transactionType string, amount float64) error {
	return nil
}

func (sm *serviceMock) GetBalance(accountID int64) (float64, error) {
	return 0, nil
}

func (sm *serviceMock) GetTransactions(accountID int64) ([]*account.Transaction, error) {
	return []*account.Transaction{}, nil
}

func TestCreateAccount(t *testing.T) {

	sm := &serviceMock{}

	f := CreateAccount(sm)

	buf := &bytes.Buffer{}
	payload := map[string]string{"name": "AA"}
	err := json.NewEncoder(buf).Encode(payload)
	assert.NoError(t, err)
	r, err := http.NewRequest("POST", "/accounts", buf)
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	f(w, r, nil)

	resp := &struct {
		ID int64 `json:"id"`
	}{}

	body, err := io.ReadAll(w.Body)
	assert.NoError(t, err)

	err = json.Unmarshal(body, resp)
	assert.NoError(t, err)

	assert.Equal(t, int64(0), resp.ID)
}

func TestProcessTransaction(t *testing.T) {

	sm := &serviceMock{}

	f := ProcessTransaction(sm)

	buf := &bytes.Buffer{}
	payload := map[string]any{"date": "2025-04-20T15:04:05Z", "type": "WIDTHDRAWAL", "amount": 7}
	err := json.NewEncoder(buf).Encode(payload)
	assert.NoError(t, err)
	r, err := http.NewRequest("POST", "/accounts/0/transactions", buf)
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	f(w, r, httprouter.Params{httprouter.Param{Key: "accountID", Value: "0"}})

	assert.Equal(t, w.Code, http.StatusOK)
}

func TestGetBalance(t *testing.T) {

	sm := &serviceMock{}

	f := GetBalance(sm)

	r, err := http.NewRequest("POST", "/accounts/0/balance", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	f(w, r, httprouter.Params{httprouter.Param{Key: "accountID", Value: "0"}})

	resp := &struct {
		Balance int64 `json:"balance"`
	}{}

	body, err := io.ReadAll(w.Body)
	assert.NoError(t, err)

	err = json.Unmarshal(body, resp)
	assert.NoError(t, err)

	assert.Equal(t, int64(0), resp.Balance)
}

func TestGetTransactionHistory(t *testing.T) {

	sm := &serviceMock{}

	f := GetTransactionHistory(sm)

	r, err := http.NewRequest("POST", "/accounts/0/transactionHistory", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	f(w, r, httprouter.Params{httprouter.Param{Key: "accountID", Value: "0"}})

	assert.Equal(t, w.Code, http.StatusOK)
}
