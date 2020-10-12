package interfaces

import (
	"context"

	"github.com/Flyewzz/tester/models"
)

type JWTManager interface {
	GenerateToken(ctx context.Context, user *models.User) (string, error)
	GetUser(ctx context.Context, token string) (*models.User, error)
}
