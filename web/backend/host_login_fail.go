package backend

import (
	"net/http"
	"html/template"
)

func Fail(w http.ResponseWriter, r *http.Request) {
	failTemplate, _ := template.ParseFiles("frontend/fail.html")
	failTemplate.Execute(w, nil)
}