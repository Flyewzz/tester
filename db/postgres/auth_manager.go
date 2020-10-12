package postgres

import (
	"context"
	"database/sql"

	"github.com/Flyewzz/tester/models"
)

type AuthManager struct {
	DB *sql.DB
}

func (this AuthManager) Authenticate(ctx context.Context, login, password string) (*models.User, error) {
	var user models.User
	err := this.DB.QueryRow(
		`SELECT id, nick, email, name, password 
		FROM users 
		WHERE (nick = $1 OR email = $1)
		AND password = $2`, login, password).Scan(
		&user.ID, &user.Nickname,
		&user.Email, &user.Name, &user.Password)
	return &user, err
}

func (this AuthManager) SignUp(ctx context.Context, user *models.User) error {
	_, err := this.DB.Exec(
		`INSERT INTO users (nick, email, name, password)
		VALUES ($1, $2, $3, $4)`, user.Nickname, user.Email, user.Name, user.Password)
	return err
}
