package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"

	_ "github.com/lib/pq"

	"github.com/dyhalmeida/golang-order/internal/infra/database"
	"github.com/dyhalmeida/golang-order/internal/usecase"
	"github.com/dyhalmeida/golang-order/pkg/rabbitmq"
)

func main() {
	db, err := sql.Open("postgres", "postgresql://postgres:postgres@localhost:5432/orders?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	orderRepository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPrice(orderRepository)

	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	msgRabbitmqChannel := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, msgRabbitmqChannel)
	rabbitmqWorker(msgRabbitmqChannel, uc)

}

func rabbitmqWorker(msgChan chan amqp.Delivery, uc *usecase.CalculateFinalPrice) {
	fmt.Println("Starting rabbitmq")
	for msg := range msgChan {
		var input usecase.OrderInput
		err := json.Unmarshal(msg.Body, &input)
		if err != nil {
			panic(err)
		}
		output, err := uc.Execute(input)
		if err != nil {
			panic(err)
		}
		msg.Ack(false)
		fmt.Println("Mensagem processada e salva no banco:", output)
	}
}
