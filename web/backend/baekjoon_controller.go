package backend

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func Echo_Baekjoon_Index(c echo.Context) error {
	var baekjoon_index = "SELECT no, title FROM baekjoon_solution ORDER BY no ASC"
	result := BaekjoonSelectQuery(db1, baekjoon_index)

	return c.Render(http.StatusOK, "baekjoon.html", result)
}

func Echo_Baekjoon_Content_View(c echo.Context) error {
	resno := c.FormValue("No")
	var select_string = "SELECT no, title, content FROM baekjoon_solution WHERE no = " + "'" + resno + "';"
	result := BaekjoonContentQuery(db1, select_string)

	return c.Render(http.StatusOK, "baekjoon_contents.html", result)
}

func Echo_Baekjoon_Write_View(c echo.Context) error {
	if c.Request().Method == "POST" {
		dt := time.Now()
		f_dt := dt.Format("2006-01-02")

		resno := c.FormValue("no")
		restitle := c.FormValue("title")
		rescontent := c.FormValue("ir1")

		var insert_string = "INSERT INTO baekjoon_solution (No, Title, Writer, Content, Date, Click) VALUES (" + "'" + resno + "'" + ", " + "'" + restitle + "'" + ", '김광호', " + "'" + rescontent + "'" + ", " + "'" + f_dt + "'" + ", " + "0" + ");"
		InsertQuery(db1, insert_string)
		http.Redirect(c.Response(), c.Request(), "/menu/?Handler=b_main", http.StatusFound)
	} else {
		return c.Render(http.StatusOK, "baekjoon_write.html", 0)
	}
	return c.String(0, "ERROR")
}
