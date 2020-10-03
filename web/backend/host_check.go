package backend

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func Echo_Login_Check(c echo.Context) error {
	resid := c.FormValue("id")
	respw := c.FormValue("pw")
	/*sess, _ := session.Get("session", c)
	sess.Values["id"] = resid
	sess.Values["pw"] = respw
	sess.Save(c.Request(), c.Response())
	log.Println("session : ", sess)
	*/

	//cookie := new(http.Cookie)
	cookie := &http.Cookie {
		Name: c.FormValue("id"),
		Value: c.FormValue("pw"),
		Path: "/",
		SameSite: http.SameSiteNoneMode,
		Secure: true,
		HttpOnly: false,
		Expires: time.Now().Add(24 * time.Hour),
	}
	c.Request().AddCookie(cookie)
	c.SetCookie(cookie)
	//http.SetCookie(c.Response(), cookie)
	/*cookie.Name = c.FormValue("id")
	cookie.Value = c.FormValue("pw")
	cookie.Path = "/"
	cookie.SameSite = 4
	cookie.Secure = true
	cookie.HttpOnly = false
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)*/

	if (resid == id) && (respw == pw) {
		fmt.Println("OK")
		c.Redirect(http.StatusMovedPermanently, "/main/")
		//http.Redirect(c.Response(), c.Request(), "/main/", http.StatusFound)
	} else {
		fmt.Println("U R Not Admin !")
		c.Redirect(http.StatusMovedPermanently, "/fail/")
		//http.Redirect(c.Response(), c.Request(), "/fail/", http.StatusFound)
	}
	return c.String(0, "ERROR")
}

func Login_Check(w http.ResponseWriter, r *http.Request) {
	//ip 주소 검색
	/*addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String() + "\n")
			}
		}
	}*/
	
	// 아이디 비밀번호 확인
	r.ParseForm()
	if r.Method == "POST" {
		resid := r.FormValue("id")
		respw := r.FormValue("pw")
		if (resid == id) && (respw == pw) {
			fmt.Println("OK")
			http.Redirect(w, r, "/main/", http.StatusFound)
		} else {
			fmt.Println("U R Not Admin !")
			http.Redirect(w, r, "/fail/", http.StatusFound)
		}
	}
}