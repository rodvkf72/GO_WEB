package backend

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

//EchoHostIndex is the first screen of the admin page
func EchoHostIndex(c echo.Context) error {
	hostcookie, _ := c.Cookie("KKH")
	if hostcookie != nil {
		fmt.Println("OK")
		c.Redirect(http.StatusMovedPermanently, "/main/")
	} else {
		//return c.Render(http.StatusOK, "index.html", "0")
	}
	return c.Render(http.StatusOK, "index.html", "0")
}

//EchoUserIndex is the first screen of the user page
func EchoUserIndex(c echo.Context) error {
	/*
		sess, _ := session.Get("session", c)
		sess.Save(c.Request(), c.Response())
		log.Println("n session : ", sess.Values["id"])
	*/

	//쿠키가 있는지 없는지 체크하는 코드
	/*
		cresult := readCookie(c, "KKH")	//cn = cookie name
		if cresult != "normal" {
			return c.String(http.StatusOK, "cookie error")
		}
	*/

	/*
		cookie, err := c.Cookie("KKH")
		if err != nil {
			return err
		}
		fmt.Println(cookie.Name)
		fmt.Println(cookie.Value)
	*/
	return c.Render(http.StatusOK, "main.html", 0)
}
