package backend

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

/*
쿠키를 생성하는 코드. cn은 쿠키명, cv는 쿠키값에 사용된다.
*/
func writeCookie(c echo.Context, cn string, cv string) {
	cookie := new(http.Cookie)
	cookie.Name = cn
	cookie.Value = cv
	cookie.Path = "/"
	cookie.SameSite = http.SameSiteLaxMode //http.SameSiteNoneMode
	cookie.Secure = false
	cookie.HttpOnly = false
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
}

/*
쿠키를 읽는 코드. (추후 사용 예정)
*/
func readCookie(c echo.Context, cn string) string {
	cookie, err := c.Cookie(cn)
	if err != nil {
		return "cookie error"
	}
	fmt.Println(cookie.Name)
	fmt.Println(cookie.Value)
	return "normal"
}
