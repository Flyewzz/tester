package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func ConfigureHandlers(r *mux.Router, api *ApiManager) {
	r.HandleFunc("/", api.MainHandler).Methods("GET")
	r.HandleFunc("/test/{id}", api.TaskCheckerHandler).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(viper.GetString("static_path"))))
	// r.HandleFunc("/requests/{id}", uh.RequestIdHandler).Methods("GET")
}
