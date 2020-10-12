package interfaces

import (
	"context"

	"github.com/Flyewzz/tester/models"
)

type SessionManager interface {
	Save(ctx context.Context, session string, user *models.User) error
	GetUser(ctx context.Context, session string) (*models.User, error)
	Remove(ctx context.Context, session string) error
}
