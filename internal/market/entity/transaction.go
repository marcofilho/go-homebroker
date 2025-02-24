package entity

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID           string    `json:"id"`
	SellingOrder *Order    `json:"selling_order"`
	BuyingOrder  *Order    `json:"buying_order"`
	Shares       int       `json:"shares"`
	Price        float64   `json:"price"`
	Total        float64   `json:"total"`
	DateTime     time.Time `json:"date_time"`
}

func NewTransaction(sellingOrder, buyingOrder *Order, shares int, price float64) *Transaction {
	return &Transaction{
		ID:           uuid.New().String(),
		SellingOrder: sellingOrder,
		BuyingOrder:  buyingOrder,
		Shares:       shares,
		Price:        price,
		Total:        0.0,
		DateTime:     time.Now(),
	}
}

func (t *Transaction) Process() {
	processor := NewOrderProcessor(t)
	processor.Process()
}
