package transaction

import (
	"github.com/gabrielcamurcab/planejador-financeiro-go/model/transaction"
)

// MockTransactionRepository é uma implementação de um repositório de transações falso para uso em testes.
type MockTransactionRepository struct {
	Transactions []*transaction.Transaction
}

// InsertTransaction insere uma nova transação no repositório falso.
func (m *MockTransactionRepository) InsertTransaction(transaction *transaction.Transaction) error {
	m.Transactions = append(m.Transactions, transaction)
	return nil
}

// GetTransactions retorna todas as transações no repositório falso.
func (m *MockTransactionRepository) GetTransactions() ([]*transaction.Transaction, error) {
	return m.Transactions, nil
}

// GetPositiveTransactions retorna todas as transações positivas no repositório falso.
func (m *MockTransactionRepository) GetPositiveTransactions() ([]*transaction.Transaction, error) {
	var positiveTransactions []*transaction.Transaction
	for _, t := range m.Transactions {
		if t.Type == 1 {
			positiveTransactions = append(positiveTransactions, t)
		}
	}
	return positiveTransactions, nil
}

// GetNegativeTransactions retorna todas as transações negativas no repositório falso.
func (m *MockTransactionRepository) GetNegativeTransactions() ([]*transaction.Transaction, error) {
	var negativeTransactions []*transaction.Transaction
	for _, t := range m.Transactions {
		if t.Type == 0 {
			negativeTransactions = append(negativeTransactions, t)
		}
	}
	return negativeTransactions, nil
}

// GetTransactionById retorna uma transação específica do repositório falso com base no ID.
func (m *MockTransactionRepository) GetTransactionById(id int) (*transaction.Transaction, error) {
	for _, t := range m.Transactions {
		if t.ID == id {
			return t, nil
		}
	}
	return nil, nil
}

// NewMockTransactionRepository cria uma nova instância de MockTransactionRepository.
func NewMockTransactionRepository() *MockTransactionRepository {
	return &MockTransactionRepository{}
}
