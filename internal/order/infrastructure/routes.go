package infrastructure

import "net/http"

func RegisterRoutes(mux *http.ServeMux, controller *OrderController) {
	mux.HandleFunc("POST /orders", controller.Create)
}
