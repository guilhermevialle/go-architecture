package user

import (
	"net/http"

	"github.com/guilhermevialle/go-architecture/internal/user/application"
	"github.com/guilhermevialle/go-architecture/internal/user/infrastructure"
)

func NewModule(mux *http.ServeMux) {
	repo := infrastructure.NewInMemoryUserRepository()
	createUserUC := application.NewCreateUser(repo)
	httpHandler := infrastructure.NewUserHttpHandler(createUserUC)

	infrastructure.RegisterRoutes(mux, httpHandler)
}
