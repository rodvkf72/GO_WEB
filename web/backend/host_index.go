package backend

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"html/template"
)

/*
관리자의 첫 화면.
 */
func Echo_Host_Index(c echo.Context) error {
	hostcookie, _:= c.Cookie("KKH")
	if hostcookie != nil {
		fmt.Println("OK")
		c.Redirect(http.StatusMovedPermanently, "/main/")
	} else {
		//return c.Render(http.StatusOK, "index.html", "0")
	}
	return c.Render(http.StatusOK, "index.html", "0")
}

//이전 방식
func Host_Index(w http.ResponseWriter, r *http.Request) {
	indexTemplate, _ := template.ParseFiles("frontend/index.html")
	indexTemplate.Execute(w, nil)
}