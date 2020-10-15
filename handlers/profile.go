package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Flyewzz/tester/models"
)

func (api *ApiManager) ProfileHandler(w http.ResponseWriter, r *http.Request) {
	props := r.Context().Value("props").(map[string]interface{})
	user := props["user"].(*models.User)
	attempts, err := api.TaskManager.GetUserAttempts(user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(attempts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(data))
}
