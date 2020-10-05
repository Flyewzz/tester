package main

import (
	"fmt"
	"net/http"

	"github.com/Flyewzz/tester/handlers"
	"github.com/spf13/viper"
)

func main() {
	PrepareConfig()
	r := NewRouter()
	c := CorsSetup()
	corsHandler := c.Handler(r)
	apiManager := PrepareApiManager()
	handlers.ConfigureHandlers(r, apiManager)
	fmt.Println("Server is listening...")
	http.ListenAndServe(":"+viper.GetString("port"), corsHandler)
}
