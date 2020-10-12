package interfaces

import (
	"context"

	"github.com/Flyewzz/tester/models"
)

type AuthManager interface {
	Authenticate(ctx context.Context, login, password string) (*models.User, error)
	SignUp(ctx context.Context, user *models.User) error
}
