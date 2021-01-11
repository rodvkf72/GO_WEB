package backend

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

//EchoBaekjoonIndex is shows the coding problem
func EchoBaekjoonIndex(c echo.Context) error {
	var baekjoonindex = "SELECT no, title FROM baekjoon_solution ORDER BY no ASC"
	result := BaekjoonSelectQuery(db1, baekjoonindex, "index")

	hostcookie, _ := c.Cookie("KKH")
	if hostcookie != nil {
		//result.Cookie = "TRUE"
		result = append(result, baekjooncontentview{Cookie: "TRUE"})
	} else {
		//result.Cookie = "FALSE"
		result = append(result, baekjooncontentview{Cookie: "FALSE"})
	}

	return c.Render(http.StatusOK, "baekjoon.html", result)
}

//EchoBaekjoonContentView is shows the content of the selected problem
func EchoBaekjoonContentView(c echo.Context) error {
	resno := c.FormValue("No")
	var selectstring = "SELECT no, title, content FROM baekjoon_solution WHERE no = " + "'" + resno + "';"
	result := BaekjoonSelectQuery(db1, selectstring, "content")

	return c.Render(http.StatusOK, "baekjoon_contents.html", result)
}

//EchoBaekjoonWriteView is shows the window to write a coding problem
func EchoBaekjoonWriteView(c echo.Context) error {
	if c.Request().Method == "POST" {
		dt := time.Now()
		fdt := dt.Format("2006-01-02")

		resno := c.FormValue("no")
		restitle := c.FormValue("title")
		rescontent := c.FormValue("ir1")

		var insertstring = "INSERT INTO baekjoon_solution (No, Title, Writer, Content, Date, Click) VALUES (" + "'" + resno + "'" + ", " + "'" + restitle + "'" + ", '김광호', " + "'" + rescontent + "'" + ", " + "'" + fdt + "'" + ", " + "0" + ");"
		InsertQuery(db1, insertstring)
		http.Redirect(c.Response(), c.Request(), "/menu/?Handler=b_main", http.StatusFound)
	} else {
		return c.Render(http.StatusOK, "baekjoon_write.html", 0)
	}
	return c.String(0, "ERROR")
}

//EchoBaekjoonSearch is problem search function
func EchoBaekjoonSearch(c echo.Context) error {
	resno := c.FormValue("no")

	query := "SELECT no, title FROM baekjoon_solution WHERE no = " + "'" + resno + "'" + ";"
	result := BaekjoonSelectQuery(db1, query, "index")

	return c.Render(http.StatusOK, "baekjoon.html", result)
}
