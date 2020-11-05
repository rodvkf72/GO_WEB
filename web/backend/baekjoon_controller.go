package backend

import (
	"github.com/labstack/echo"
	"net/http"
)

func Echo_Baekjoon_Index(c echo.Context) error {
	//var baekjoon_index = "SELECT no, title FROM baekjoon_solution ORDER BY no ASC"
	return c.Render(http.StatusOK, "baekjoon.html", 0)
}

func Echo_Baekjoon_Content_View(c echo.Context) error {
	return c.Render(http.StatusOK, "baekjoon.html", 0)
}

func Echo_Baekjoon_Write_View(c echo.Context) error {
	return c.Render(http.StatusOK, "baekjoon.html", 0)
}