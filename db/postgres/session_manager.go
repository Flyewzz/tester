package postgres

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Flyewzz/tester/models"
	"github.com/go-redis/redis/v8"
)

type SessionManager struct {
	DB *redis.Client
}

func (this SessionManager) Save(ctx context.Context, session string, user *models.User) error {
	duration, ok := ctx.Value("duration").(int)
	if !ok {
		duration = 0
	}
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return this.DB.Set(ctx, session, data,
		time.Duration(duration)*time.Second).Err()
}

func (this SessionManager) GetUser(ctx context.Context, session string) (*models.User, error) {
	data, err := this.DB.Get(ctx, session).Result()
	if err != nil {
		return nil, err
	}
	var user models.User
	err = json.Unmarshal([]byte(data), &user)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
func (this SessionManager) Remove(ctx context.Context, session string) error {
	return this.DB.Del(ctx, session).Err()
}
