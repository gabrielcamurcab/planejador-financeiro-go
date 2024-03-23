package transaction

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gabrielcamurcab/planejador-financeiro-go/model/transaction"
	"github.com/gabrielcamurcab/planejador-financeiro-go/util"
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	var transanctions = transaction.Transactions{
		transaction.Transaction{
			Title:     "Salário",
			Amount:    1440.40,
			Type:      0,
			CreatedAt: util.StringToTime("2024-03-22T21:33:00"),
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
