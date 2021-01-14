package backend

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

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
	f, err := os.OpenFile("iplog.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				log.Println(ipnet.IP.String())
			}
		}
	}
	log.Println("-----------------")

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
