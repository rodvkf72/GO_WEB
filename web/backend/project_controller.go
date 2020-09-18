package backend

import (
	"github.com/labstack/echo"
	"html/template"
	"net/http"
)

func Echo_Project_Index(c echo.Context) error {
	return c.Render(http.StatusOK, "project.html", 0)
}

func Echo_Project_Content_View(c echo.Context) error {
	c.Request().ParseForm()
	resno := c.Request().FormValue("No")
	result := []project_view{}
	if resno != "" {
		result = return_contents(resno)
	}
	return c.Render(http.StatusOK, "project_contents_view.html", result)
}

func Echo_Project_Write_View(c echo.Context) error {
	return c.Render(http.StatusOK, "project_write.html", 0)
}

func Project_Index(w http.ResponseWriter, r *http.Request) {
	var projectlist string = "frontend/project_list.html"
	var projectcontents string = "frontend/project_contents.html"
	p_indexTemplate, _ := template.ParseFiles("frontend/project.html", header, footer, projectlist, projectcontents, leftside)
	p_indexTemplate.Execute(w, nil)
}

func Project_Content_View(w http.ResponseWriter, r *http.Request) {
	p_contentTemplate, _ := template.ParseFiles("frontend/project_contents_view.html", header, footer, leftside)
	r.ParseForm()
	resno := r.FormValue("No")
	result := []project_view{}
	if resno != "" {
		result = return_contents(resno)
	}
	p_contentTemplate.Execute(w, result)
}

func Project_Write_View(w http.ResponseWriter, r *http.Request) {
	p_writeTemplate, _ := template.ParseFiles("frontend/project_write.html", header, footer, leftside)
	p_writeTemplate.Execute(w, nil)
}
