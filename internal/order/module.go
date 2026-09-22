package order

import (
	"net/http"

	"github.com/guilhermevialle/go-architecture/internal/order/application"
	"github.com/guilhermevialle/go-architecture/internal/order/infrastructure"
)

func NewModule(mux *http.ServeMux) {
	repo := infrastructure.NewInMemoryOrderRepository()
	createOrderUC := application.NewCreateOrder(repo)
	httpHandler := infrastructure.NewOrderHttpHandler(createOrderUC)

	infrastructure.RegisterRoutes(mux, httpHandler)
}
