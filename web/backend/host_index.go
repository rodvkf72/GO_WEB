package backend

import (
	"net/http"
	"html/template"
)

func Host_Index(w http.ResponseWriter, r *http.Request) {
	indexTemplate, _ := template.ParseFiles("frontend/index.html")
	indexTemplate.Execute(w, nil)