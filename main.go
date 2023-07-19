package main

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	_ "github.com/lib/pq"

	"github.com/dyhalmeida/golang-order/internal/infra/database"
	"github.com/dyhalmeida/golang-order/internal/usecase"
)

func main() {
	db, err := sql.Open("postgres", "postgresql://postgres:postgres@localhost:5432/orders?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	orderRepository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPrice(orderRepository)

	input := usecase.OrderInput{
		ID:    uuid.New().String(),
		Price: 10.0,
		Tax:   1.0,
	}

	output, err := uc.Execute(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}
