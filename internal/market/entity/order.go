package entity

type Order struct {
	ID string
	Investor *Investor
	Asset *Asset
	Shares int 
	PendingShares int
	Price   float64
	OrderType string
	Status string
	Transactions []*Transaction
}

func NewOrder (orderID string, Investor*Investor, asset *Asset, shares int, price float64, Ordertype string,) *Order{
	return &Order{
        ID: 			orderID,
        Investor:		Investor,
        Asset:			asset,
        Shares:			shares,
        Price:			price,
        OrderType:		Ordertype,
        Status: 		"OPEN",
        Transactions:	[]*Transaction{},
    }
}