package backend

import (
	"net/http"
	"html/template"
)

func Project_Index(w http.ResponseWriter, r *http.Request) {
	p_indexTemplate, _ := template.ParseFiles("frontend/project.html", header, footer)
	p_indexTemplate.Execute(w, nil)
}

func Project_Content_View(w http.ResponseWriter, r *http.Request) {
	p_contentTemplate, _ := template.ParseFiles("frontend/project_contents.html", header, footer)
	r.ParseForm()
	resno := r.FormValue("No")
	result := []project_view{}
	if resno != "" {
		result = return_contents(resno)
	}
	p_contentTemplate.Execute(w, result)
}

func Project_Write_View(w http.ResponseWriter, r *http.Request) {
	p_writeTemplate, _ := template.ParseFiles("frontend/project_write.html", header, footer)
	p_writeTemplate.Execute(w, nil)
}