package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Flyewzz/tester/models"
	"github.com/gorilla/mux"
)

func (api *ApiManager) LoginHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(1024)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ctx := context.Background()
	params := r.FormValue

	login := params("login")
	password := params("password")

	user, err := api.AuthManager.Authenticate(ctx, login, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	token, err := api.JWTManager.GenerateToken(ctx, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	cookie := &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(time.Duration(720) * time.Hour),
		Path:    "/",
	}
	http.SetCookie(w, cookie)
	r.AddCookie(cookie)
	w.Header().Add("Content-Type", "text/plain")
	w.Write([]byte(token))
}

func (api *ApiManager) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	params := r.PostFormValue
	nick := params("nick")
	email := params("email")
	name := params("name")
	password := params("password")
	//! Don't forget to add a validator!

	err = api.AuthManager.SignUp(context.Background(), &models.User{
		Nickname: nick,
		Email:    email,
		Name:     name,
		Password: password,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	w.WriteHeader(201)
}

func (api ApiManager) AuthMiddleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			fmt.Println("Malformed token")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		token := authHeader[1]
		ctx := context.Background()
		// user, err := api.SessionManager.GetUser(context.TODO(), token)
		user, err := api.JWTManager.GetUser(ctx, token)
		if err != nil || user == nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		vars := mux.Vars(r)
		ctx = context.WithValue(ctx, "props", map[string]interface{}{
			"vars": vars,
			"user": user,
		})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
