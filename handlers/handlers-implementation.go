package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/Flyewzz/tester/checker"
	"github.com/spf13/viper"
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
	err := r.ParseMultipartForm(0)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	params := r.FormValue
	code := params("code")
	fmt.Println(code)
	// copy example
	prog, _ := os.Create("checker/task/main.cpp")
	defer prog.Close()
	_, err = prog.WriteString(code)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	program := checker.NewCppProgram(
		viper.GetString("programPath"),
		"100m",
		"110m",
		".500",
		800,
	)
	verdicts := program.Check()
	data, _ := json.Marshal(verdicts)
	w.Write(data)
}
