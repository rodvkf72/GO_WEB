package backend

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

/*
 문제 풀이 선택 화면
*/
func Echo_Baekjoon_Index(c echo.Context) error {
	var baekjoon_index = "SELECT no, title FROM baekjoon_solution ORDER BY no ASC"
	result := BaekjoonSelectQuery(db1, baekjoon_index)

	hostcookie, _ := c.Cookie("KKH")
	if hostcookie != nil {
		//result.Cookie = "TRUE"
		result = append(result, baekjoon_view{Cookie: "TRUE"})
	} else {
		//result.Cookie = "FALSE"
		result = append(result, baekjoon_view{Cookie: "FALSE"})
	}

	return c.Render(http.StatusOK, "baekjoon.html", result)
}

/*
 문제 풀이 내용 화면
*/
func Echo_Baekjoon_Content_View(c echo.Context) error {
	resno := c.FormValue("No")
	var select_string = "SELECT no, title, content FROM baekjoon_solution WHERE no = " + "'" + resno + "';"
	result := BaekjoonContentQuery(db1, select_string)

	return c.Render(http.StatusOK, "baekjoon_contents.html", result)
}

/*
 문제 풀이 작성 화면
*/
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

/*
 데이터베이스에 있는 문제 검색 기능
*/
func Echo_Baekjoon_Search(c echo.Context) error {
	resno := c.FormValue("no")

	query := "SELECT no, title FROM baekjoon_solution WHERE no = " + "'" + resno + "'" + ";"
	result := BaekjoonSelectQuery(db1, query)

	return c.Render(http.StatusOK, "baekjoon.html", result)
}
