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
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

func (api *ApiManager) TaskInfoGetHandler(w http.ResponseWriter, r *http.Request) {
	strId := r.URL.Query().Get("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(w, "Incorrect id", http.StatusInternalServerError)
		return
	}
	fmt.Printf("id: %d\n", id)

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
	vars := mux.Vars(r)

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

	program := checker.NewCppProgram(
		programPath,
		"100m",
		"110m",
		".500",
		800,
	)

	tests := api.TestLoader.Load(id)
	verdicts := program.Check(tests)
	data, _ := json.Marshal(verdicts)
	w.Write(data)
}
