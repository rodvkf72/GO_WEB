package backend

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/labstack/echo"
)

/*
관리자의 첫 화면.
*/
func Echo_Host_Index(c echo.Context) error {
	hostcookie, _ := c.Cookie("KKH")
	if hostcookie != nil {
		fmt.Println("OK")
		c.Redirect(http.StatusMovedPermanently, "/main/")
	} else {
		//return c.Render(http.StatusOK, "index.html", "0")
	}
	return c.Render(http.StatusOK, "index.html", "0")
}

//이전 방식
func Host_Index(w http.ResponseWriter, r *http.Request) {
	indexTemplate, _ := template.ParseFiles("frontend/index.html")
	indexTemplate.Execute(w, nil)
}

/*
일반 사용자가 보게 될 첫 화면
*/
func Echo_User_Index(c echo.Context) error {
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

// 일반 사용자가 보게 될 첫 화면
func User_Index(w http.ResponseWriter, r *http.Request) {
	mainTemplate, _ := template.ParseFiles("frontend/main.html", header, footer, leftside)
	mainTemplate.Execute(w, nil)

	//r.ParseForm()
	//Get 파라미터 및 정보 출력
	//fmt.Println("default : ", r.Form)
	//fmt.Println("path : ", r.URL.Path)
	//fmt.Println("param : ", r.Form["test_param"])
	//Parameter 전체 출력
	//for k, v := range r.Form {
	//	fmt.Println("key : ", k)
	//	fmt.Println("val : ", strings.Join(v, ""))
	//}
	//기본 출력
	//fmt.Fprintf(w, "Golang WebServer Working!")
}
