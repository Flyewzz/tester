package handlers

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func (api *ApiManager) MainHandler(w http.ResponseWriter, r *http.Request) {
	type Data struct {
		Title string
	}
	t, _ := template.ParseFiles("pages/index.html") // Parse template file.
	title := "Сложите два числа"
	t.Execute(w, Data{
		Title: title,
	})
}

func (api *ApiManager) TaskCheckerHandler(w http.ResponseWriter, r *http.Request) {
	code := r.PostFormValue("code")
	fmt.Println(code)
	// copy example
	prog, _ := os.Create("main.cpp")
	defer prog.Close()
	_, err := prog.WriteString(code)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("OK"))
	// io.Copy(f, file)
}
