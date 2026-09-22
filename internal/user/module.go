package user

import (
	"net/http"

	"github.com/guilhermevialle/go-architecture/internal/user/application"
	"github.com/guilhermevialle/go-architecture/internal/user/infrastructure"
)

func NewModule(mux *http.ServeMux, repo application.UserRepository) {
	createUserUC := application.NewCreateUser(repo)
	controller := infrastructure.NewUserController(createUserUC)

	infrastructure.RegisterRoutes(mux, controller)
}
