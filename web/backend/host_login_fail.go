package backend

import (
	"github.com/labstack/echo"
	"net/http"
	"html/template"
)

/*
로그인 실패 시 동작
 */
func Echo_Fail(c echo.Context) error {
	return c.Render(http.StatusOK, "fail.html", "0")
}

//이전 방식
func Fail(w http.ResponseWriter, r *http.Request) {
	failTemplate, _ := template.ParseFiles("frontend/fail.html")
	failTemplate.Execute(w, nil)
}