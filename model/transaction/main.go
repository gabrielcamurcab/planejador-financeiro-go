package transaction

type Transaction struct {
	Title  string  `json:"title"`
	Amount float32 `json:"amount"`
	Type   int     `json:"type"`
}

type Transactions []Transaction
