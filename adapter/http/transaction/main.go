package transaction

import (
	"encoding/json"
	"net/http"

	"github.com/gabrielcamurcab/planejador-financeiro-go/repository/transaction"
)

type TransactionHandler struct {
	repo *transaction.TransactionRepository
}

// Criar um construtor diretamente para TransactionHandler que aceita TransactionRepository
func NewTransactionHandler(repo *transaction.TransactionRepository) *TransactionHandler {
	return &TransactionHandler{repo: repo}
}

func (handler *TransactionHandler) GetTransactions(w http.ResponseWriter, r *http.Request) {
	transactions, err := handler.repo.GetTransactions()
	if err != nil {
		http.Error(w, "Erro ao buscar transações", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if transactions != nil {
		json.NewEncoder(w).Encode(transactions)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "Nenhuma transação encontrada"})
	}
}

func (handler *TransactionHandler) CreateATransaction(w http.ResponseWriter, r *http.Request) {
	var t transaction.Transaction
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Erro ao decodificar o JSON da transação", http.StatusBadRequest)
		return
	}

	err = handler.repo.InsertTransaction(&t)
	if err != nil {
		http.Error(w, "Erro ao criar a transação", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Transação cadastrada com sucesso!"})
}
