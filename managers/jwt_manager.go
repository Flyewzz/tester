package managers

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/Flyewzz/tester/models"
	"github.com/dgrijalva/jwt-go"
)

type JWTManager struct {
	duration  time.Duration
	secretKey []byte
}

func NewJWTManager(duration time.Duration, secretKey string) *JWTManager {
	return &JWTManager{
		duration:  duration,
		secretKey: []byte(secretKey),
	}
}

func (this JWTManager) GenerateToken(ctx context.Context, user *models.User) (string, error) {
	// We are happy with the credentials, so build a token. We've given it
	// an expiry of 1 hour.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":   time.Now().Add(this.duration).Unix(),
		"iat":   time.Now().Unix(),
		"id":    user.ID,
		"nick":  user.Nickname,
		"name":  user.Name,
		"email": user.Email,
		"role":  "user",
	})
	//! Fix role (to add new roles)
	tokenString, err := token.SignedString(this.secretKey)
	if err != nil {
		return "", errors.New("Token generation failed")
	}
	fmt.Printf(`token: %s\n`, tokenString)
	return tokenString, nil
}

func (this JWTManager) GetUser(ctx context.Context, token string) (*models.User, error) {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	unsignedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return this.secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := unsignedToken.Claims.(jwt.MapClaims); ok && unsignedToken.Valid {
		id, err := strconv.Atoi(fmt.Sprintf("%v", claims["id"]))
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		user := &models.User{
			ID:       id,
			Nickname: fmt.Sprintf("%v", claims["nick"]),
			Name:     fmt.Sprintf("%v", claims["name"]),
			Email:    fmt.Sprintf("%v", claims["email"]),
		}
		fmt.Printf("A token was received:\n- nick: %s\n- name: %s\n", claims["nick"], claims["name"])
		return user, nil
	} else {
		fmt.Println(err)
		return nil, err
	}
}
