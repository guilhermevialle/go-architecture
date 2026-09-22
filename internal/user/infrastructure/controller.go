package infrastructure

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/guilhermevialle/go-architecture/internal/user/application"
)

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserController struct {
	createUser *application.CreateUser
}

func NewUserController(createUser *application.CreateUser) *UserController {
	return &UserController{createUser: createUser}
}

func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := c.createUser.Execute(r.Context(), req.Name, req.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		slog.Error("failed to encode response", "error", err)
	}
}
