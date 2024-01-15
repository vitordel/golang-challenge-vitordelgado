package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Message struct {
	ID        string    `json:"id"`
	Cliente   string    `json:"cliente"`
	Timestamp time.Time `json:"timestamp"`
	Payload   string    `json:"payload"`
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
		"messages", // Nome da fila
		false,      // Durable
		false,      // Delete when unused
		false,      // Exclusive
		false,      // No-wait
		nil,        // Arguments
	)
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := ch.Consume(
		q.Name, // Fila
		"",     // Consumidor
		true,   // Auto-ack
		false,  // Exclusive
		false,  // No-local
		false,  // No-wait
		nil,    // Args
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Server aguardando mensagens...")

	for msg := range msgs {
		var message Message
		err := json.Unmarshal(msg.Body, &message)
		if err != nil {
			log.Println("Erro ao decodificar mensagem:", err)
		} else {
			fmt.Printf("Mensagem recebida: %+v\n", message)
		}
	}
}
