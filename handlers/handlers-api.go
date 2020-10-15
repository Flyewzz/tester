package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func ConfigureHandlers(r *mux.Router, api *ApiManager) {
	r.HandleFunc("/", api.AuthMiddleware(http.HandlerFunc(api.TaskInfoGetHandler))).Methods("GET")
	r.HandleFunc("/test/{id}", api.AuthMiddleware(http.HandlerFunc(api.TaskCheckerHandler))).Methods("POST")
	r.HandleFunc("/profile/attempts", api.AuthMiddleware(http.HandlerFunc(api.ProfileHandler))).Methods("GET")
	authRouter := r.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/login", api.LoginHandler).Methods("POST")
	authRouter.HandleFunc("/signup", api.SignUpHandler).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(viper.GetString("static_path"))))

}
