package entity

type Order struct {
	ID            string         `json:"id"`
	Investor      *Investor      `json:"investor"`
	Asset         *Asset         `json:"asset"`
	Shares        int            `json:"shares"`
	PendingShares int            `json:"pending_shares"`
	Price         float64        `json:"price"`
	OrderType     string         `json:"order_type"`
	Status        string         `json:"status"`
	Transactions  []*Transaction `json:"transaction"`
}

func NewOrder(orderID string, investor *Investor, asset *Asset, shares int, price float64, orderType string) *Order {
	return &Order{
		ID:            orderID,
		Investor:      investor,
		Asset:         asset,
		Shares:        shares,
		PendingShares: shares,
		Price:         price,
		OrderType:     orderType,
		Status:        "OPEN",
		Transactions:  []*Transaction{},
	}
}

func (o *Order) ApplyTrade(tradedShares int) {
	if tradedShares > o.PendingShares {
		tradedShares = o.PendingShares
	}

	o.PendingShares -= tradedShares

	if o.PendingShares == 0 {
		o.Status = "CLOSED"
	}
}

func (o *Order) AddTransaction(transaction *Transaction) {
	o.Transactions = append(o.Transactions, transaction)
}
