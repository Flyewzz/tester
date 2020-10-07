package handlers

import (
	"github.com/Flyewzz/tester/interfaces"
)

type ApiManager struct {
	TestLoader interfaces.TestLoader
}

func NewApiManager(loader interfaces.TestLoader) *ApiManager {
	return &ApiManager{
		TestLoader: loader,
	}
}
