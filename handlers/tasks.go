package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Flyewzz/tester/checker"
	"github.com/Flyewzz/tester/models"
)

func (api *ApiManager) TaskInfoGetHandler(w http.ResponseWriter, r *http.Request) {
	strId := r.URL.Query().Get("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(w, "Incorrect id", http.StatusInternalServerError)
		return
	}

	taskInfo, taskCount, err := api.TaskStorage.GetInfo(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	type Answer struct {
		ID          int    `json:"id"`
		Text        string `json:"text"`
		Ram         string `json:"ram"`
		HDD         string `json:"hdd"`
		Time        string `json:"time"`
		Samples     string `json:"samples"`
		Limitations string `json:"limitations"`
		TaskCount   int    `json:"task_count"`
	}
	answer := Answer{
		ID:          taskInfo.ID,
		Text:        taskInfo.Text,
		Ram:         taskInfo.Ram,
		HDD:         taskInfo.HDD,
		Time:        taskInfo.Time,
		Samples:     taskInfo.Samples,
		Limitations: taskInfo.Limitations,
		TaskCount:   taskCount,
	}
	data, err := json.Marshal(answer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(data))
}

func (api *ApiManager) TaskCheckerHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	err := r.ParseMultipartForm(1024 * 1024)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	props := r.Context().Value("props").(map[string]interface{})
	vars := props["vars"].(map[string]string)

	strId, ok := vars["id"]
	if !ok {
		http.Error(w, "Id is required", http.StatusInternalServerError)
		return
	}
	id, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(w, "Incorrect id", http.StatusInternalServerError)
		return
	}
	params := r.FormValue
	code := params("code")

	taskInfo, _, err := api.TaskStorage.GetInfo(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	timeLimit, err := strconv.Atoi(taskInfo.Time)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Deviation depends on computer's power
	program, err := api.ProgramManager.Create(taskInfo.Ram,
		taskInfo.HDD,
		".800",
		timeLimit+api.Deviation, code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tests := api.TestLoader.Load(id)

	user := props["user"].(*models.User)
	verdict := program.Check(tests)
	program.Remove()
	_, err = api.TaskManager.SetStatus(user.ID, taskInfo.ID, verdict.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data, _ := json.Marshal([]*checker.Verdict{
		verdict,
	})
	w.Write(data)
}
