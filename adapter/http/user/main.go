package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gabrielcamurcab/planejador-financeiro-go/model/user"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user user.User
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Erro ao decodificar dados do usuário", http.StatusBadRequest)
		return
	}

	user.UUID = uuid.NewString()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Erro ao gerar o hash da senha", http.StatusInternalServerError)
		log.Printf("Erro ao gerar o hash da senha: %v", err)
		return
	}

	user.Password = string(hashedPassword)

	err = sendMessageToQueue(user)
	if err != nil {
		http.Error(w, "Erro ao enviar mensagens para a fila", http.StatusInternalServerError)
		log.Printf("Erro: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Usuário criado com sucesso!"})
}

func sendMessageToQueue(user user.User) error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"create-user",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	body, err := json.Marshal(user)
	if err != nil {
		return err
	}

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		return err
	}

	return nil
}
