package infrastructure

import (
	"context"
	"sync"

	"github.com/guilhermevialle/go-architecture/internal/order/domain"
)

type InMemoryOrderRepository struct {
	mu     sync.RWMutex
	orders []*domain.Order
}

func NewInMemoryOrderRepository() *InMemoryOrderRepository {
	return &InMemoryOrderRepository{orders: make([]*domain.Order, 0)}
}

func (r *InMemoryOrderRepository) Save(ctx context.Context, order *domain.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.orders = append(r.orders, order)
	return nil
}
