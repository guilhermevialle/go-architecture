package order

import (
	"net/http"

	"github.com/guilhermevialle/go-architecture/internal/order/application"
	"github.com/guilhermevialle/go-architecture/internal/order/infrastructure"
)

func NewModule(mux *http.ServeMux) {
	repo := infrastructure.NewInMemoryOrderRepository()
	createOrderUC := application.NewCreateOrder(repo)
	controller := infrastructure.NewOrderController(createOrderUC)

	infrastructure.RegisterRoutes(mux, controller)
}
