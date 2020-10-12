package models

import (
	"github.com/dgrijalva/jwt-go"
)

// jwtCustomClaims are custom claims extending default ones.
type JWTClaims struct {
	jwt.StandardClaims
	User
}
