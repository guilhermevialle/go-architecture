package application

import (
	"context"

	"github.com/guilhermevialle/go-architecture/internal/user/domain"
)

type UserRepository interface {
	Save(ctx context.Context, user *domain.User) error
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
}
