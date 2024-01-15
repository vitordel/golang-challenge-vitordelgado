package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Message struct {
	ID        string    `json:"id"`
	Cliente   string    `json:"cliente"`
	Timestamp time.Time `json:"timestamp"`
	Payload   string    `json:"payload"`
}

type LoginResponse struct {
	Token  string `json:"token"`
	UserID string `json:"userId"`
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"messages",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Cliente enviando mensagens...")

	for {

		username := "vitor"
		password := "123456"

		userID, token, err := PerformLogin(username, password)
		if err != nil {
			log.Println("Error during login:", err)
			return
		}
		message := Message{
			ID:        uuid.New().String(),
			Cliente:   userID,
			Timestamp: time.Now(),
			Payload:   token,
		}

		body, err := json.Marshal(message)
		if err != nil {
			log.Println("Erro ao codificar mensagem:", err)
			continue
		}

		err = ch.Publish(
			"",
			q.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        body,
			},
		)
		if err != nil {
			log.Println("Erro ao publicar mensagem:", err)
		}

		time.Sleep(5 * time.Second)
	}
}

func PerformLogin(username, password string) (string, string, error) {
	loginURL := "http://localhost:8080/login"

	loginData := map[string]string{"username": username, "password": password}
	loginBody, err := json.Marshal(loginData)
	if err != nil {
		log.Println("Error encoding login request:", err)
		return "", "", err
	}

	resp, err := http.Post(loginURL, "application/json", bytes.NewBuffer(loginBody))
	if err != nil {
		log.Println("Error performing login request:", err)
		return "", "", err
	}
	defer resp.Body.Close()

	var loginResponse LoginResponse
	err = json.NewDecoder(resp.Body).Decode(&loginResponse)
	if err != nil {
		log.Println("Error decoding login response:", err)
		return "", "", err
	}

	return loginResponse.UserID, loginResponse.Token, nil
}
