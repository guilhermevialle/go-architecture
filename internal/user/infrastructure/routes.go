package infrastructure

import "net/http"

func RegisterRoutes(mux *http.ServeMux, controller *UserController) {
	mux.HandleFunc("POST /users", controller.Create)
}
