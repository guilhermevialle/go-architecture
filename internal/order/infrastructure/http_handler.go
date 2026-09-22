package infrastructure

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/guilhermevialle/go-architecture/internal/order/application"
)

type CreateOrderRequest struct {
	UserID string `json:"userID"`
	Price  int    `json:"price"`
}

type OrderHttpHandler struct {
	createOrder *application.CreateOrder
}

func NewOrderHttpHandler(createOrder *application.CreateOrder) *OrderHttpHandler {
	return &OrderHttpHandler{createOrder: createOrder}
}

func (c *OrderHttpHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order, err := c.createOrder.Execute(r.Context(), req.UserID, req.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(order); err != nil {
		slog.Error("failed to encode response", "error", err)
	}
}
