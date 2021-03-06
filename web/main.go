package main

import (
	"GO_WEB/web/backend"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/crypto/acme/autocert"
)

//Template is structure to use template function
/*
 template 기능을 사용하기 위한 구조체
*/
type Template struct {
	templates *template.Template
}

//GetTempFilesFromFolders is scans file path
func GetTempFilesFromFolders(folders []string) []string {
	var filepaths []string
	for _, folder := range folders {
		files, _ := ioutil.ReadDir(folder)

		for _, file := range files {
			if strings.Contains(file.Name(), ".html") {
				filepaths = append(filepaths, folder+file.Name())
			}
		}
	}
	return filepaths
}

//Render is function to use template function
/*
 template 기능을 사용하기 위한 함수
 w : http.status
 name : html 명
 data : 전달하고자 하는 데이터
*/
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	dirs := []string{"./frontend/",
		"./frontend/static/include/",
		"./frontend/baekjoon/",
		"./frontend/error/",
		"./frontend/game/",
		"./frontend/notice_board/",
		"./frontend/project/"}

	tempfiles := GetTempFilesFromFolders(dirs)
	t := &Template{
		//templates: template.Must(template.ParseGlob("./frontend/*.html")),
		templates: template.Must(template.ParseFiles(tempfiles...)),
	}

	fs := http.FileServer(http.Dir("./frontend/static")) //파일 서버를 지정

	e := echo.New()
	e.AutoTLSManager.Cache = autocert.DirCache("/static/ssl/") //TLS의 캐시 위치를 지정
	e.Use(middleware.Logger())                                 //미들웨어에서 로거를 사용
	e.Use(middleware.Recover())                                //미들웨어에서 복구를 사용
	e.Static("/static/", "public")
	e.Renderer = t

	e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", fs)))
	e.POST("/static/*", echo.WrapHandler(http.StripPrefix("/static/", fs)))
	e.GET("/", backend.EchoHostIndex)
	e.POST("/check/", backend.EchoLoginCheck)
	e.GET("/main/", backend.EchoUserIndex)
	e.GET("/fail/", backend.EchoFail)

	//main에서는 하나로 처리하지만 컨트롤러에 의해 global_info.go와 handler.go 파일의 함수에서 여러 개로 처리
	e.GET("/menu/*", backend.EchoRequestHandler)

	/*
		기존 코드에서는 GET, POST 방식 구분 없이 main에서 사용이 가능했기에 핸들러 함수 하나로 처리가 가능하였으나 Echo 에서는 시작할 때
		GET, POST 방식을 지정하므로 아래와 같이 데이터베이스에 직접적으로 데이터가 전송되는 부분은 POST 방식으로 처리
	*/
	e.POST("/menu/g_write", backend.EchoGameWriteView)        //POST 처리되어 있어서 게시글 작성 창 진입을 관리자가 아닌 다른 사람이 했는지 쿠키로 체크 할 필요가 없음.
	e.POST("/menu/n_write", backend.EchoNoticeboardWriteView) //html에서 관리자 쿠키가 없으면 버튼 자체가 뜨지 않게 해 놓았음
	e.POST("/menu/b_write", backend.EchoBaekjoonWriteView)
	e.POST("/single_img_upload/", backend.SingleImgUpload)
	e.POST("/multi_img_upload/", backend.MultiImgUpload)
	e.POST("/codemirror", backend.CodeAjax)

	//tls.LoadX509KeyPair("./frontend/static/ssl/private.crt", "./frontend/static/ssl/private.key")
	//e.Logger.Fatal(e.StartTLS(":433", "./frontend/static/ssl/private.crt", "./frontend/static/ssl/private.key")) //https 보안연결.

	err := e.Start(":9090")
	if err != nil {
		e.Logger.Fatal(err)
	}

}
