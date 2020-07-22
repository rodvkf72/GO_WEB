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
	Project_View := project_view{}
	Project_Views := []project_view{}
	if resno == "1" {
		Project_View.No = resno
		Project_View.Title = "공공데이터 파싱"
		Project_View.Content = "공공데이터 파싱을 활용하였다." 
		Project_View.Root = "/static/image/main_2.png"
		Project_Views = append(Project_Views, Project_View)
	}
	p_contentTemplate.Execute(w, Project_Views)
}

func Project_Write_View(w http.ResponseWriter, r *http.Request) {
	p_writeTemplate, _ := template.ParseFiles("frontend/project_write.html", header, footer)
	p_writeTemplate.Execute(w, nil)
}