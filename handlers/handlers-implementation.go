package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"text/template"
)

func (api *ApiManager) MainHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Задание: Правильное сложение чисел"))

	type Data struct {
		Title string
	}
	t, _ := template.ParseFiles("pages/index.html") // Parse template file.
	title := "Сложите два числа"
	t.Execute(w, Data{
		Title: title,
	}) // merge.
	// method, url := r.URL.Query().Get("method"), r.URL.Query().Get("url")
	// if !CheckMethodValid(method) {
	// 	http.Error(w, "Bad request", http.StatusBadRequest)
	// 	return
	// }
	// task := &models.Task{
	// 	Method: method,
	// 	Url:    url,
	// }
	// resCh := hd.Dispatcher.AddNewTask(task)
	// result := <-resCh
	// if result.Error != nil {
	// 	http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
	// 	log.Println(result.Error)
	// 	return
	// }
	// data, err := json.Marshal(result.Response)
	// if err != nil {
	// 	http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
	// 	log.Println(err)
	// 	return
	// }
	// w.Write(data)
}

func (api *ApiManager) TaskCheckerHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// r.ParseMultipartForm(5 * 1024 * 1024)
	file, _, err := r.FormFile("file")
	if err != nil {
		log.Println(err)
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer file.Close()
	// copy example
	f, err := os.OpenFile("./downloaded", os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	io.Copy(f, file)
	// method, url := r.URL.Query().Get("method"), r.URL.Query().Get("url")

}
