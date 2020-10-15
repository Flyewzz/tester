package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Flyewzz/tester/db"
	"github.com/Flyewzz/tester/handlers"
	"github.com/spf13/viper"
)

func main() {
	PrepareConfig()
	r := NewRouter()
	c := CorsSetup()

	corsHandler := c.Handler(r)
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
	defer db.Close()
	db.SetMaxOpenConns(viper.GetInt("db.max_connections"))
	apiManager := PrepareApiManager(db)
	handlers.ConfigureHandlers(r, apiManager)
	fmt.Println("Server is listening...")
	http.ListenAndServe(":"+viper.GetString("port"), corsHandler)
}
