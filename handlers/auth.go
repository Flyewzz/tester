package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Flyewzz/tester/models"
	"github.com/Flyewzz/tester/validators"
	"github.com/gorilla/mux"
)

func (api *ApiManager) LoginHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(1024)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	user := &models.User{}
	err = validators.LoginUserValidator(r, user)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	ctx := context.Background()
	user, err = api.AuthManager.Authenticate(ctx, user.Login, user.Password)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
	token, err := api.JWTManager.GenerateToken(ctx, user)
	if err != nil {
		http.Error(w, err.Error(), 500)
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
	err := r.ParseMultipartForm(1024)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	user := &models.User{}
	err = validators.SignUpUserValidator(r, user)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	ctx := context.Background()
	err = api.AuthManager.SignUp(ctx, user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	user, err = api.AuthManager.Authenticate(ctx, user.Login, user.Password)
	if err != nil || user == nil {
		http.Error(w, err.Error(), 500)
		return
	}

	token, err := api.JWTManager.GenerateToken(ctx, user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(201)
	w.Write([]byte(token))
}

func (api ApiManager) AuthMiddleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			fmt.Println("Malformed token")
			w.WriteHeader(401)
			return
		}
		token := authHeader[1]
		ctx := context.Background()
		// user, err := api.SessionManager.GetUser(context.TODO(), token)
		user, err := api.JWTManager.GetUser(ctx, token)
		if err != nil || user == nil {
			http.Error(w, err.Error(), 401)
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
