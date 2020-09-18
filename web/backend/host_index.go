package backend

import (
	"github.com/labstack/echo"
	"net/http"
	"html/template"
)

func Echo_Host_Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", "0")
}

func Host_Index(w http.ResponseWriter, r *http.Request) {
	indexTemplate, _ := template.ParseFiles("frontend/index.html")
	indexTemplate.Execute(w, nil)
}