package transaction

import (
	"database/sql"
	"time"
)

type Transaction struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Amount    float32   `json:"amount"`
	Type      int       `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (repo *TransactionRepository) InsertTransaction(t *Transaction) error {
	query := `INSERT INTO transactions (title, amount, type, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5)`

	_, err := repo.db.Exec(query, t.Title, t.Amount, t.Type, time.Now(), time.Now())
	if err != nil {
		return err
	}

	return nil
}

func (repo *TransactionRepository) GetTransactions() ([]*Transaction, error) {
	var transactions []*Transaction

	rows, err := repo.db.Query("SELECT * FROM transactions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		t := &Transaction{}
		err := rows.Scan(&t.ID, &t.Title, &t.Amount, &t.Type, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}
