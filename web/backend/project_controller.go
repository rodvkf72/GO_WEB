package backend

import (
	"net/http"

	"github.com/labstack/echo"
)

//EchoProjectIndex is first page of the project tab
/*
프로젝트 탭의 첫 화면
*/
func EchoProjectIndex(c echo.Context) error {
	return c.Render(http.StatusOK, "project.html", 0)
}

//EchoProjectContentView is shows project detail content
/*
프로젝트의 자세한 내용을 출력하기 위한 함수
*/
func EchoProjectContentView(c echo.Context) error {
	c.Request().ParseForm()
	resno := c.Request().FormValue("No")
	result := []projectview{}
	if resno != "" {
		result = returncontents(resno)
	}
	return c.Render(http.StatusOK, "project_contents_view.html", result)
}

//EchoProjectWriteView is not yet used... this is project content add function
/*
프로젝트 내용을 추가하기 위한 함수 (아직 사용되지 않음)
*/
func EchoProjectWriteView(c echo.Context) error {
	return c.Render(http.StatusOK, "project_write.html", 0)
}
