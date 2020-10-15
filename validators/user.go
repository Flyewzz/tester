package validators

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Flyewzz/tester/models"
	"github.com/thedevsaddam/govalidator"
)

func SignUpUserValidator(r *http.Request, u *models.User) error {
	rules := govalidator.MapData{
		"login":    []string{"required", "between:6,25"},
		"email":    []string{"required", "between:4,80"},
		"name":     []string{"required", "between:3,80"},
		"password": []string{"required", "between:6,30"},
	}
	messages := govalidator.MapData{
		"login": []string{
			"required:Логин не может быть пустым",
			"between:Логин должен быть длиной от 6 до 25 символов"},
		"email": []string{
			"required:Электронная почта должна быть заполнена",
			"between:Некорректная электронная почта"},
		"name": []string{
			"required:Имя не может быть пустым",
			"between:Некорректный ввод имени"},
		"password": []string{
			"required:Пароль не может быть пустым",
			"between:Пароль должен быть длиной от 6 до 30 символов"},
	}

	opts := govalidator.Options{
		Request:         r,        // request object
		Rules:           rules,    // rules map
		Messages:        messages, // custom message map (Optional)
		Data:            u,
		RequiredDefault: true, // all the field to be pass the rules
	}
	v := govalidator.New(opts)
	e := v.Validate()
	if len(e) == 0 {
		u.Name = r.FormValue("name")
		u.Password = r.FormValue("password")
		u.Email = r.FormValue("email")
		u.Login = r.FormValue("login")
		return nil
	}
	err := map[string]interface{}{"validationError": e}
	data, _ := json.Marshal(err)
	return errors.New(string(data))
}

func LoginUserValidator(r *http.Request, u *models.User) error {
	rules := govalidator.MapData{
		"login":    []string{"required", "between:6,25"},
		"password": []string{"required", "between:6,30"},
	}
	messages := govalidator.MapData{
		"login": []string{
			"required:Логин не может быть пустым",
			"between:Логин должен быть длиной от 6 до 25 символов"},
		"password": []string{
			"required:Пароль не может быть пустым",
			"between:Пароль должен быть длиной от 6 до 30 символов"},
	}

	opts := govalidator.Options{
		Request:  r,        // request object
		Rules:    rules,    // rules map
		Messages: messages, // custom message map (Optional)
		Data:     u,
	}
	v := govalidator.New(opts)
	e := v.Validate()
	if len(e) == 0 {
		u.Login = r.FormValue("login")
		u.Password = r.FormValue("password")
		return nil
	}
	err := map[string]interface{}{"validationError": e}
	data, _ := json.Marshal(err)
	return errors.New(string(data))
}
