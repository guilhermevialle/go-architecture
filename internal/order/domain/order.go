package domain

type Order struct {
	UserID string
	Price  int
}

func NewOrder(userID string, price int) (*Order, error) {
	return &Order{
		UserID: userID,
		Price:  price,
	}, nil
}
