package transaction_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gabrielcamurcab/planejador-financeiro-go/repository/transaction"
	"github.com/gabrielcamurcab/planejador-financeiro-go/util"
)

// Testes

func TestInsertTransaction(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Erro ao criar mock de banco de dados: %v", err)
	}
	defer db.Close()

	repo := transaction.NewTransactionRepository(db)

	mock.ExpectExec("INSERT INTO transactions").
		WithArgs("Transaction 1", 100.00, 1, time.Now(), time.Now()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.InsertTransaction(&transaction.Transaction{
		Title:     "Transaction 1",
		Amount:    100.0,
		Type:      1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		t.Errorf("Erro ao inserir transação: %v", err)
	}
}

func TestGetTransactions(t *testing.T) {
	// Cria um objeto de mock para o banco de dados
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Erro ao criar mock do banco de dados: %v", err)
	}
	defer db.Close()

	// Cria o repositório com o banco de dados mock
	repo := transaction.NewTransactionRepository(db)

	// Define o comportamento esperado para a consulta no banco de dados
	rows := sqlmock.NewRows([]string{"id", "title", "amount", "type", "created_at", "updated_at"}).
		AddRow(1, "Transaction 1", 100.00, 0, util.StringToTime("2024-03-30 12:00:00"), util.StringToTime("2024-03-30 12:00:00")).
		AddRow(2, "Transaction 2", 50.00, 1, util.StringToTime("2024-03-30 12:00:00"), util.StringToTime("2024-03-30 12:00:00"))

	// Configura o mock para corresponder à consulta esperada
	mock.ExpectQuery("SELECT \\* FROM transactions").WillReturnRows(rows)

	// Chama a função que você deseja testar
	transactions, err := repo.GetTransactions()
	if err != nil {
		t.Errorf("Erro ao buscar transações: %v", err)
	}

	// Verifica se as transações foram retornadas corretamente
	if len(transactions) != 2 {
		t.Errorf("Número incorreto de transações retornadas. Esperado: 2, Recebido: %d", len(transactions))
	}
}

func TestGetPositiveTransactions(t *testing.T) {
	// Cria um objeto de mock para o banco de dados
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Erro ao criar mock do banco de dados: %v", err)
	}
	defer db.Close()

	// Cria o repositório com o banco de dados mock
	repo := transaction.NewTransactionRepository(db)

	// Define o comportamento esperado para a consulta no banco de dados
	rows := sqlmock.NewRows([]string{"id", "title", "amount", "type", "created_at", "updated_at"}).
		AddRow(1, "Transaction 1", 100.00, 1, util.StringToTime("2024-03-30T12:00:00"), util.StringToTime("2024-03-30T12:00:00"))

	// Configura o mock para corresponder à consulta esperada
	mock.ExpectQuery("SELECT id, title, amount, type, created_at, updated_at FROM transactions WHERE type = 1").WillReturnRows(rows)

	// Chama a função que você deseja testar
	transactions, err := repo.GetPositiveTransactions()
	if err != nil {
		t.Errorf("Erro ao buscar transações: %v", err)
	}

	// Verifica se as transações foram retornadas corretamente
	if len(transactions) != 1 {
		t.Errorf("Número incorreto de transações retornadas. Esperado: 1, Recebido: %d", len(transactions))
	}
}

func TestGetNegativeTransactions(t *testing.T) {
	// Cria um objeto de mock para o banco de dados
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Erro ao criar mock do banco de dados: %v", err)
	}
	defer db.Close()

	// Cria o repositório com o banco de dados mock
	repo := transaction.NewTransactionRepository(db)

	// Define o comportamento esperado para a consulta no banco de dados
	rows := sqlmock.NewRows([]string{"id", "title", "amount", "type", "created_at", "updated_at"}).
		AddRow(1, "Transaction 1", 100.00, 0, util.StringToTime("2024-03-30T12:00:00"), util.StringToTime("2024-03-30T12:00:00"))

	// Configura o mock para corresponder à consulta esperada
	mock.ExpectQuery("SELECT id, title, amount, type, created_at, updated_at FROM transactions WHERE type = 0").WillReturnRows(rows)

	// Chama a função que você deseja testar
	transactions, err := repo.GetNegativeTransactions()
	if err != nil {
		t.Errorf("Erro ao buscar transações: %v", err)
	}

	// Verifica se as transações foram retornadas corretamente
	if len(transactions) != 1 {
		t.Errorf("Número incorreto de transações retornadas. Esperado: 1, Recebido: %d", len(transactions))
	}
}

func TestGetTransactionById(t *testing.T) {
	// Cria um objeto de mock para o banco de dados
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Erro ao criar mock do banco de dados: %v", err)
	}
	defer db.Close()

	// Cria o repositório com o banco de dados mock
	repo := transaction.NewTransactionRepository(db)

	// Define o comportamento esperado para a consulta no banco de dados
	rows := sqlmock.NewRows([]string{"id", "title", "amount", "type", "created_at", "updated_at"}).
		AddRow(1, "Transaction 1", 100.00, 0, util.StringToTime("2024-03-30 12:00:00"), util.StringToTime("2024-03-30 12:00:00"))

	mock.ExpectQuery("SELECT id, title, amount, type, created_at, updated_at FROM transactions WHERE id = ?").
		WillReturnRows(rows)

	// Chama a função que você deseja testar
	transaction, err := repo.GetTransactionById(1)
	if err != nil {
		t.Errorf("Erro ao buscar transações: %v", err)
	}

	// Verifica se as transações foram retornadas corretamente
	if transaction == nil {
		t.Error("Nenhuma transação retornada.")
	}
}
