package application

import (
	"context"
	"fmt"

	"github.com/guilhermevialle/go-architecture/internal/order/domain"
)

type CreateOrder struct {
	repo OrderRepository
}

func NewCreateOrder(repo OrderRepository) *CreateOrder {
	return &CreateOrder{repo: repo}
}

func (uc *CreateOrder) Execute(ctx context.Context, userID string, price int) (*domain.Order, error) {
	order, err := domain.NewOrder(userID, price)
	if err != nil {
		return nil, err
	}

	if err := uc.repo.Save(ctx, order); err != nil {
		return nil, fmt.Errorf("failed to save order: %w", err)
	}

	return order, nil
}
