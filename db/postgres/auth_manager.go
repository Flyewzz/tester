package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/Flyewzz/tester/models"
	"golang.org/x/crypto/bcrypt"
)

type AuthManager struct {
	DB      *sql.DB
	Timeout time.Duration
}

func (this AuthManager) Authenticate(ctx context.Context, login, password string) (*models.User, error) {
	var user models.User
	var hashedPassword []byte

	ctx, cancel := context.WithTimeout(ctx, this.Timeout)
	defer cancel()
	err := this.DB.QueryRowContext(
		ctx,
		`SELECT id, login, email, name, password 
		FROM users 
		WHERE (login = $1 OR email = $1)`, login).Scan(
		&user.ID, &user.Login,
		&user.Email, &user.Name, &hashedPassword)
	if err != nil {
		return nil, err
	}
	if err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password)); err != nil {
		return nil, errors.New("Invalid login credentials")
	}

	return &user, err
}

func (this AuthManager) SignUp(ctx context.Context, user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	password := string(hashedPassword)
	ctx, cancel := context.WithTimeout(ctx, this.Timeout)
	defer cancel()
	_, err = this.DB.ExecContext(
		ctx,
		`INSERT INTO users (login, email, name, password)
		VALUES ($1, $2, $3, $4)`, user.Login, user.Email, user.Name, password)
	return err
}
