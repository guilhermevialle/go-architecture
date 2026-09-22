package infrastructure

import "net/http"

func RegisterRoutes(mux *http.ServeMux, httpHandler *OrderHttpHandler) {
	mux.HandleFunc("POST /orders", httpHandler.Create)
}
