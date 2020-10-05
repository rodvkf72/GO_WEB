package backend

import (
	"github.com/labstack/echo"
	"net/http"
	"html/template"
)

/*
관리자의 첫 화면.
 */
func Echo_Host_Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", "0")
}

//이전 방식
func Host_Index(w http.ResponseWriter, r *http.Request) {
	indexTemplate, _ := template.ParseFiles("frontend/index.html")
	indexTemplate.Execute(w, nil)
}