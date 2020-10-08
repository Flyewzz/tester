package handlers

import (
	"github.com/Flyewzz/tester/interfaces"
)

type ApiManager struct {
	TestLoader  interfaces.TestLoader
	TaskStorage interfaces.TaskStorage
	Deviation   int
}

func NewApiManager(loader interfaces.TestLoader,
	storage interfaces.TaskStorage, deviation int) *ApiManager {
	return &ApiManager{
		TestLoader:  loader,
		TaskStorage: storage,
		Deviation:   deviation,
	}
}
