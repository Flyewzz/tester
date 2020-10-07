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

func (api *ApiManager) MainHandler(w http.ResponseWriter, r *http.Request) {
	// type Data struct {
	// 	Title string
	// 	Task  string
	// }
	// t, _ := template.ParseFiles("pages/index.html") // Parse template file.
	// title := "На вход подается массив чисел через пробел. <br> Выведите все отрицательные числа в виде " +
	// 	"отсортированного по возрастанию массива."
	// task := ""
	// t.Execute(w, Data{
	// 	Title: title,
	// 	Task:  task,
	// })
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
	fmt.Printf("id: %d\n", id)
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
