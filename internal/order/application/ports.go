package application

import (
	"context"

	"github.com/guilhermevialle/go-architecture/internal/order/domain"
)

type OrderRepository interface {
	Save(ctx context.Context, order *domain.Order) error
}
