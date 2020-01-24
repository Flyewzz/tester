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
	apiManager := PrepareApiManager()
	handlers.ConfigureHandlers(r, apiManager)
	fmt.Println("Server is listening...")
	http.ListenAndServe(":"+viper.GetString("port"), r)
}
