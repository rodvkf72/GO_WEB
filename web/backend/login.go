package backend

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

//EchoLoginCheck is manager login check function
/*
관리자가 로그인 시 아이디와 비밀번호를 체크하는 기능이다.
KKH라는 이름의 쿠키를 생성한다.
로그인 성공 시 /main/으로 이동하며 실패 시 /fail/로 이동한다.
*/
func EchoLoginCheck(c echo.Context) error {
	resid := c.FormValue("id")
	respw := c.FormValue("pw")

	//쿠키 이전 세션으로 처리하려고 시도했던 것
	/*sess, _ := session.Get("session", c)
	sess.Values["id"] = resid
	sess.Values["pw"] = respw
	sess.Save(c.Request(), c.Response())
	log.Println("session : ", sess)
	*/

	//http 쿠키 사용 방법
	/*"cookie := &http.Cookie {
		Name: c.FormValue("id"),
		Value: c.FormValue("pw"),
		Path: "/",
		SameSite: http.SameSiteNoneMode,
		Secure: true,
		HttpOnly: false,
		Expires: time.Now().Add(24 * time.Hour),
	}*/

	//c.Request().AddCookie(cookie)

	if (resid == id) && (respw == pw) {
		fmt.Println("OK")
		cookie := new(http.Cookie)
		cookie.Name = "KKH"
		cookie.Value = "Blog"
		cookie.Path = "/"
		cookie.SameSite = http.SameSiteLaxMode
		cookie.Secure = false
		cookie.HttpOnly = false
		cookie.Expires = time.Now().Add(24 * time.Hour)
		c.SetCookie(cookie)
		c.Redirect(http.StatusMovedPermanently, "/main/")
		//http.Redirect(c.Response(), c.Request(), "/main/", http.StatusFound)
	} else {
		fmt.Println("U R Not Admin !")
		c.Redirect(http.StatusMovedPermanently, "/fail/")
		//http.Redirect(c.Response(), c.Request(), "/fail/", http.StatusFound)
	}
	return c.String(0, "ERROR")
}

//ip 체크. 단, 사용자의 ip 이외에 연결되는 ip 모두 출력됨
/*
func Ip_Check(c echo.Context) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				log.Println(ipnet.IP.String() + "\n")
			}
		}
	}
}
*/

//EchoFail is works when login fails
/*
로그인 실패 시 동작
*/
func EchoFail(c echo.Context) error {
	return c.Render(http.StatusOK, "fail.html", "0")
}
