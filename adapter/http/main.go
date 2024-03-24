package http

import (
	"database/sql"
	"net/http"

	"github.com/gabrielcamurcab/planejador-financeiro-go/adapter/http/actuator"
	TransactionHandler "github.com/gabrielcamurcab/planejador-financeiro-go/adapter/http/transaction"
	"github.com/gabrielcamurcab/planejador-financeiro-go/repository/transaction"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Init(db *sql.DB) {
	repo := transaction.NewTransactionRepository(db)
	handler := TransactionHandler.NewTransactionHandler(repo)
	r := mux.NewRouter()

	r.HandleFunc("/transactions", handler.GetTransactions).Methods("GET")
	r.HandleFunc("/transactions/positive", handler.GetPositiveTransactions).Methods("GET")
	r.HandleFunc("/transactions/negative", handler.GetNegativeTransactions).Methods("GET")

	r.HandleFunc("/transactions", handler.CreateATransaction).Methods("POST")

	r.HandleFunc("/health", actuator.Health).Methods("GET")

	r.Handle("/metrics", promhttp.Handler()).Methods("GET")

	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)
}
