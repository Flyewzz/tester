package handlers

import (
	"github.com/Flyewzz/tester/interfaces"
)

type ApiManager struct {
	TestLoader  interfaces.TestLoader
	TaskStorage interfaces.TaskStorage
	AuthManager interfaces.AuthManager
	JWTManager  interfaces.JWTManager
	Deviation   int
}
