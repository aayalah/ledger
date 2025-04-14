package main

import (
	"fmt"
	"github/ledger/config"
	"github/ledger/handlers"
	"github/ledger/service"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	service := service.New()
	config, err := config.New()
	if err != nil {
		log.Fatalf("error creating configuration")
	}

	router := httprouter.New()
	router.POST("/accounts", handlers.CreateAccount(service))
	router.POST("/accounts/:accountID/transactions", handlers.ProcessTransaction(service))
	router.GET("/accounts/:accountID/balance", handlers.GetBalance(service))
	router.GET("/accounts/:accountID/transactionHistory", handlers.GetTransactionHistory(service))

	http.ListenAndServe(fmt.Sprintf(":%v", config.PORT), router)
}
