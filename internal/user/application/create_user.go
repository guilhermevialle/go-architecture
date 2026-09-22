package application

import (
	"context"
	"fmt"

	"github.com/guilhermevialle/go-architecture/internal/user/domain"
)

type CreateUser struct {
	repo UserRepository
}

func NewCreateUser(repo UserRepository) *CreateUser {
	return &CreateUser{repo: repo}
}

func (uc *CreateUser) Execute(ctx context.Context, name, email string) (*domain.User, error) {
	user, err := domain.NewUser(name, email)
	if err != nil {
		return nil, err
	}

	if err := uc.repo.Save(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to save user: %w", err)
	}

	return user, nil
}
