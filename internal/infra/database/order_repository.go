package database

import (
	"database/sql"

	"github.com/dyhalmeida/golang-order/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db}
}

func (orderRepository OrderRepository) Save(order *entity.Order) error {
	_, err := orderRepository.Db.Exec("INSERT INTO orders (id, price, tax, final_price) VALUES (?,?,?,?)",
		order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (orderRepository *OrderRepository) GetTotalTransactions() (int, error) {
	var total int
	err := orderRepository.Db.QueryRow("SELECT COUNT(*) FROM orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}
