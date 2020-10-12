package main

import (
	"log"
	"os"
	"time"

	"github.com/Flyewzz/tester/checker"
	"github.com/Flyewzz/tester/db"
	"github.com/Flyewzz/tester/db/postgres"
	"github.com/Flyewzz/tester/db/sqlite"
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

func PrepareApiManager() *handlers.ApiManager {
	db, err := db.ConnectToDB(viper.GetString("db.host"),
		viper.GetString("db.user"), viper.GetString("db.password"),
		viper.GetString("db.database"), viper.GetInt("db.port"))
	if err != nil {
		log.Fatalf("Error with database: %v\n", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error with connection to the database: %v\n", err)
	}
	db.SetMaxOpenConns(viper.GetInt("db.max_connections"))

	loader := checker.TestLoader{
		Path: viper.GetString("task_path"),
	}
	taskStorage := sqlite.NewTaskStorage(viper.GetString("sqlite.path"))
	deviation := viper.GetInt("time.execution.deviation")

	authManager := postgres.AuthManager{
		DB: db,
	}
	jwtManager := managers.NewJWTManager(
		time.Hour*time.Duration(viper.GetInt("jwt.duration")),
		viper.GetString("jwt.secret_key"))

	return &handlers.ApiManager{
		TestLoader:  loader,
		TaskStorage: taskStorage,
		Deviation:   deviation,
		AuthManager: authManager,
		JWTManager:  jwtManager,
	}
}
