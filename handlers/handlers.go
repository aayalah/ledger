package handlers

import (
	"encoding/json"
	"github/ledger/account"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

const (
	ACCOUNT_ID_PARAM = "accountID"
)

type service interface {
	CreateAccount(name string) int64
	RecordTransaction(accountID int64, date time.Time, transactionType string, amount float64) error
	GetBalance(accountID int64) (float64, error)
	GetTransactions(accountID int64) ([]*account.Transaction, error)
}

func CreateAccount(s service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading create account body: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		account := &struct {
			Name string `json:"name"`
		}{}

		err = json.Unmarshal(body, account)
		if err != nil {
			log.Printf("Error unmarshalling account body: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		id := s.CreateAccount(account.Name)

		respBodyStr := struct {
			ID int64 `json:"id"`
		}{
			ID: id,
		}

		respBody, err := json.Marshal(respBodyStr)
		if err != nil {
			log.Printf("Error marshalling created account response: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(respBody)
	}
}

func ProcessTransaction(s service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading create account body: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		accountIDStr := p.ByName(ACCOUNT_ID_PARAM)

		accountID, err := strconv.ParseInt(accountIDStr, 0, 64)
		if err != nil {
			log.Printf("Error parsing account id: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		transaction := &struct {
			Date            string  `json:"date"`
			TransactionType string  `json:"type"`
			Amount          float64 `json:"amount"`
		}{}

		err = json.Unmarshal(body, transaction)
		if err != nil {
			log.Printf("Error unmarshalling transaction body: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		t, err := time.Parse(time.RFC3339, transaction.Date)
		if err != nil {
			log.Printf("Error parsing date: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = s.RecordTransaction(accountID, t, transaction.TransactionType, transaction.Amount)
		if err != nil {
			log.Printf("Error processing transaction: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func GetBalance(s service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		accountIDStr := p.ByName(ACCOUNT_ID_PARAM)

		accountID, err := strconv.ParseInt(accountIDStr, 0, 64)
		if err != nil {
			log.Printf("Error parsing acount id: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		balance, err := s.GetBalance(accountID)
		if err != nil {
			log.Printf("Error getting balance: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		respBodyStr := &struct {
			Balance float64 `json:"balance"`
		}{
			Balance: balance,
		}

		respBody, err := json.Marshal(respBodyStr)
		if err != nil {
			log.Printf("Error marshalling balance response body: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(respBody)
	}
}

type transactionResp struct {
	Date            time.Time `json:"date"`
	TransactionType string    `json:"type"`
	Amount          float64   `json:"amount"`
}

func GetTransactionHistory(s service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		accountIDStr := p.ByName(ACCOUNT_ID_PARAM)

		accountID, err := strconv.ParseInt(accountIDStr, 0, 64)
		if err != nil {
			log.Printf("Error parsing account id: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		transactions, err := s.GetTransactions(accountID)
		if err != nil {
			log.Printf("Error getting transactions: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		transactionResps := make([]transactionResp, 0, len(transactions))

		for _, txn := range transactions {
			transactionResps = append(transactionResps, transactionResp{
				Date:            txn.Date,
				Amount:          txn.Amount,
				TransactionType: txn.TransactionType,
			})
		}

		respBody, err := json.Marshal(transactionResps)
		if err != nil {
			log.Printf("Error marshalling transactions response body: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(respBody)

		w.WriteHeader(http.StatusOK)
	}
}
