package infrastructure

import "net/http"

func RegisterRoutes(mux *http.ServeMux, httpHandler *UserHttpHandler) {
	mux.HandleFunc("POST /users", httpHandler.Create)
}
