package http

import (
	"net/http"

	"github.com/gabrielcamurcab/planejador-financeiro-go/adapter/http/actuator"
	"github.com/gabrielcamurcab/planejador-financeiro-go/adapter/http/transaction"
)

func Init() {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateATransaction)

	http.HandleFunc("/health", actuator.Health)

	http.ListenAndServe(":8080", nil)
}
