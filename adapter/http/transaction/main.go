package transaction

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gabrielcamurcab/planejador-financeiro-go/repository/transaction"
	"github.com/gorilla/mux"
)

type TransactionHandler struct {
	Repo *transaction.TransactionRepository
}

// Criar um construtor diretamente para TransactionHandler que aceita TransactionRepository
func NewTransactionHandler(repo *transaction.TransactionRepository) *TransactionHandler {
	return &TransactionHandler{Repo: repo}
}

func (handler *TransactionHandler) GetTransactions(w http.ResponseWriter, r *http.Request) {
	transactions, err := handler.Repo.GetTransactions()
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

	err = handler.Repo.InsertTransaction(&t)
	if err != nil {
		http.Error(w, "Erro ao criar a transação", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Transação cadastrada com sucesso!"})
}

func (handler *TransactionHandler) GetPositiveTransactions(w http.ResponseWriter, r *http.Request) {
	transactions, err := handler.Repo.GetPositiveTransactions()
	if err != nil {
		http.Error(w, "Erro ao buscar transações", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	if transactions != nil {
		json.NewEncoder(w).Encode(transactions)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "Nenhuma transação de entrada foi encontrada"})
	}
}

func (handler *TransactionHandler) GetNegativeTransactions(w http.ResponseWriter, r *http.Request) {
	transactions, err := handler.Repo.GetNegativeTransactions()
	if err != nil {
		http.Error(w, "Erro ao buscar transações", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	if transactions != nil {
		json.NewEncoder(w).Encode(transactions)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "Nenhuma transação de saída foi encontrada"})
	}
}

func (handler *TransactionHandler) GetTransactionById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	transactionId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "ID da tranasação inválido", http.StatusBadRequest)
		return
	}

	transaction, err := handler.Repo.GetTransactionById(transactionId)
	if err != nil {
		http.Error(w, "Erro ao buscar transações", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	if transaction != nil {
		json.NewEncoder(w).Encode(transaction)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "Transação não encontrada"})
	}
}
