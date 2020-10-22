package handlers

import (
	"github.com/Flyewzz/tester/interfaces"
)

type ApiManager struct {
	TestLoader     interfaces.TestLoader
	TaskStorage    interfaces.TaskStorage
	AuthManager    interfaces.AuthManager
	JWTManager     interfaces.JWTManager
	TaskManager    interfaces.TaskManager
	ProgramManager interfaces.ProgramManager
	Deviation      int
}
