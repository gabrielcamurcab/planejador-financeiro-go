package http

import (
	"database/sql"
	"net/http"

	"github.com/gabrielcamurcab/planejador-financeiro-go/adapter/http/actuator"
	TransactionHandler "github.com/gabrielcamurcab/planejador-financeiro-go/adapter/http/transaction"
	"github.com/gabrielcamurcab/planejador-financeiro-go/repository/transaction"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Init(db *sql.DB) {
	repo := transaction.NewTransactionRepository(db)
	handler := TransactionHandler.NewTransactionHandler(repo)

	http.HandleFunc("/transactions", handler.GetTransactions)

	http.HandleFunc("/transactions/create", handler.CreateATransaction)

	http.HandleFunc("/health", actuator.Health)

	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8080", nil)
}
