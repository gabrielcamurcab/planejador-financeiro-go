package transaction

import "time"

type Transaction struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Amount    float32   `json:"amount"`
	Type      int       `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Transactions []Transaction
