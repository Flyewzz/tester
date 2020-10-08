package main

import (
	"log"
	"os"

	"github.com/Flyewzz/tester/checker"
	"github.com/Flyewzz/tester/db/sqlite"
	"github.com/Flyewzz/tester/handlers"
	"github.com/spf13/viper"
)

func PrepareConfig() {
	viper.SetConfigFile(os.Args[1])
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Cannot read a config file: %v\n", err)
	}
}

func PrepareApiManager() *handlers.ApiManager {
	loader := checker.TestLoader{
		Path: viper.GetString("task_path"),
	}
	taskStorage := sqlite.NewTaskStorage(viper.GetString("sqlite.path"))
	return handlers.NewApiManager(loader, taskStorage)
}
