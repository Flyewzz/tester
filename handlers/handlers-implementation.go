package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Flyewzz/tester/checker"
	"github.com/Flyewzz/tester/models"
	uuid "github.com/satori/go.uuid"
)

func (api *ApiManager) TaskInfoGetHandler(w http.ResponseWriter, r *http.Request) {
	strId := r.URL.Query().Get("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(w, "Incorrect id", http.StatusInternalServerError)
		return
	}

	taskInfo, err := api.TaskStorage.GetInfo(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(taskInfo)
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
	err := r.ParseMultipartForm(0)
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
	fmt.Println(code)

	taskInfo, err := api.TaskStorage.GetInfo(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	folderPath := "checker/task/" +
		strings.Replace(uuid.NewV4().String(),
			"-", "", -1)

	programPath := filepath.Join(folderPath, "main.cpp")

	err = os.Mkdir(folderPath, 0700)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer os.RemoveAll(folderPath)

	programFile, err := os.Create(programPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer programFile.Close()
	_, err = programFile.WriteString(code)
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
	program := checker.NewCppProgram(
		programPath,
		taskInfo.Ram,
		taskInfo.HDD,
		".800",
		timeLimit+api.Deviation,
	)

	tests := api.TestLoader.Load(id)

	user := props["user"].(*models.User)
	verdict := program.Check(tests)
	if verdict.Status == "OK" {
		log.Printf("User %d %s solved task #%d\n", user.ID, user.Nickname, id)
	}
	data, _ := json.Marshal([]*checker.Verdict{
		verdict,
	})
	w.Write(data)
}
