package handlers

import (
	"github.com/Flyewzz/tester/interfaces"
)

type ApiManager struct {
	TestLoader  interfaces.TestLoader
	TaskStorage interfaces.TaskStorage
}

func NewApiManager(loader interfaces.TestLoader,
	storage interfaces.TaskStorage) *ApiManager {
	return &ApiManager{
		TestLoader:  loader,
		TaskStorage: storage,
	}
}
