package transaction

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gabrielcamurcab/planejador-financeiro-go/model/transaction"
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	var layout = "2006-01-02T15:04:05"
	salaryReceived, _ := time.Parse(layout, "2024-03-22T21:33:00")

	var transanctions = transaction.Transactions{
		transaction.Transaction{
			Title:     "Salário",
			Amount:    1440.40,
			Type:      0,
			CreatedAt: salaryReceived,
		},
	}

	_ = json.NewEncoder(w).Encode(transanctions)
}

func CreateATransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var res = transaction.Transactions{}
	var body, _ = io.ReadAll(r.Body)

	_ = json.Unmarshal(body, &res)

	fmt.Print(res)
}
