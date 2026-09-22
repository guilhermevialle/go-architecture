package user

import (
	"net/http"

	"github.com/guilhermevialle/go-architecture/internal/order/infrastructure"
	"github.com/guilhermevialle/go-architecture/internal/user/application"
)

func NewModule(mux *http.ServeMux) {
	repo := infrastructure.NewInMemoryUserRepository()
	createUserUC := application.NewCreateUser(repo)
	httpHandler := infrastructure.NewUserHttpHandler(createUserUC)

	infrastructure.RegisterRoutes(mux, httpHandler)
}
