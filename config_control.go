package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/Flyewzz/tester/checker"
	"github.com/Flyewzz/tester/db/postgres"
	"github.com/Flyewzz/tester/handlers"
	"github.com/Flyewzz/tester/managers"
	"github.com/spf13/viper"
)

func PrepareConfig() {
	viper.SetConfigFile(os.Args[1])
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Cannot read a config file: %v\n", err)
	}
}

func PrepareApiManager(db *sql.DB) *handlers.ApiManager {
	loader := checker.TestLoader{
		Path: viper.GetString("task_path"),
	}
	// taskStorage := sqlite.NewTaskStorage(viper.GetString("sqlite.path"))
	taskStorage := postgres.TaskStorage{
		DB: db,
	}
	deviation := viper.GetInt("time.execution.deviation")

	authManager := postgres.AuthManager{
		DB: db,
	}
	jwtManager := managers.NewJWTManager(
		time.Hour*time.Duration(viper.GetInt("jwt.duration")),
		viper.GetString("jwt.secret_key"))

	taskManager := postgres.TaskManager{
		DB: db,
	}
	return &handlers.ApiManager{
		TestLoader:  loader,
		TaskStorage: taskStorage,
		Deviation:   deviation,
		AuthManager: authManager,
		JWTManager:  jwtManager,
		TaskManager: taskManager,
	}
}
