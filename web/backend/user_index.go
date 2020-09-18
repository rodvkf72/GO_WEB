package backend

import (
	"github.com/labstack/echo"
	"net/http"
	"html/template"
)

func Echo_User_Index(c echo.Context) error {
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
