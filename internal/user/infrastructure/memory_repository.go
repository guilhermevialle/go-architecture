package infrastructure

import (
	"context"
	"sync"

	"github.com/guilhermevialle/go-architecture/internal/user/domain"
)

type InMemoryUserRepository struct {
	mu    sync.RWMutex
	users map[string]*domain.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{users: make(map[string]*domain.User)}
}

func (r *InMemoryUserRepository) Save(ctx context.Context, user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.users[user.Email] = user
	return nil
}

func (r *InMemoryUserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, ok := r.users[email]
	if !ok {
		return nil, nil
	}
	return user, nil
}
