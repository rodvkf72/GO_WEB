package main

import (
	"GO_WEB/web/backend"
	"crypto/tls"
	"html/template"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/crypto/acme/autocert"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("./frontend/*.html")),
	}
	fs := http.FileServer(http.Dir("./frontend/static")) //파일 서버를 지정

	e := echo.New()
	e.AutoTLSManager.Cache = autocert.DirCache("/static/ssl/") //TLS의 캐시 위치를 지정
	e.Use(middleware.Logger())                                 //미들웨어에서 로거를 사용
	e.Use(middleware.Recover())                                //미들웨어에서 복구를 사용
	//e.Use(session.MiddlewareWithConfig(session.Config{}))
	//e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.Static("/static/", "public")
	e.Renderer = t

	e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", fs)))
	e.POST("/static/*", echo.WrapHandler(http.StripPrefix("/static/", fs)))
	//e.POST("/static/smarteditor2/sample/photo_uploader/attach_photo.js", echo.WrapHandler(http.StripPrefix("/static/smartedit2/", fs)))
	/*e.GET("/", func (c echo.Context) error {
		store := sessions.NewCookieStore([]byte("secret"))
		sess, _ := store.Get(c.Request(), "test")
		//sess, _ := session.Get("session", c)
		sess.Options = &sessions.Options {
			Domain: "localhost",
			Path: "/",
			MaxAge: 3600 * 8,	//8 hours
			HttpOnly: true,
		}
		sess.Values["foo"] = "bar"
		sess.Save(c.Request(), c.Response())
		log.Println("sess : ", sess)
		e.Use(session.MiddlewareWithConfig(session.Config{Store: store}))
		return backend.Echo_Host_Index(c)//c.NoContent(http.StatusOK)

	store := sessions.NewCookieStore([]byte("secret"))
	store.Options = &sessions.Options {
		Domain: "localhost:9090",
		Path: "/",
		MaxAge: 3600 * 8, //8 hours
		HttpOnly: true,
	}
	e.Use(session.MiddlewareWithConfig(session.Config{Store: store}))
	})
	*/
	e.GET("/", backend.Echo_Host_Index)
	e.POST("/check/", backend.Echo_Login_Check)
	e.GET("/main/", backend.Echo_User_Index)
	e.GET("/fail/", backend.Echo_Fail)

	//main에서는 하나로 처리하지만 컨트롤러에 의해 global_info.go와 handler.go 파일의 함수에서 여러 개로 처리
	e.GET("/menu/*", backend.Echo_Request_Handler)

	/*
		기존 코드에서는 GET, POST 방식 구분 없이 main에서 사용이 가능했기에 핸들러 함수 하나로 처리가 가능하였으나 Echo 에서는 시작할 때
		GET, POST 방식을 지정하므로 아래와 같이 데이터베이스에 직접적으로 데이터가 전송되는 부분은 POST 방식으로 처리
	*/
	e.POST("/menu/g_write", backend.Echo_Game_Write_View)        //POST 처리되어 있어서 게시글 작성 창 진입을 관리자가 아닌 다른 사람이 했는지 쿠키로 체크 할 필요가 없음.
	e.POST("/menu/n_write", backend.Echo_Noticeboard_Write_View) //html에서 관리자 쿠키가 없으면 버튼 자체가 뜨지 않게 해 놓았음
	e.POST("/menu/p_write", backend.Echo_Baekjoon_Write_View)
	e.POST("/single_img_upload/", backend.SingleImgUpload)
	e.POST("/multi_img_upload/", backend.MultiImgUpload)

	/*e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "/static/main.html", echo.Map{"title" : "Page file title!!"})
	})*/
	//e.Logger.Fatal(e.StartAutoTLS(":433"))
	tls.LoadX509KeyPair("./frontend/static/ssl/private.crt", "./frontend/static/ssl/private.key")
	e.Logger.Fatal(e.StartTLS(":433", "./frontend/static/ssl/private.crt", "./frontend/static/ssl/private.key")) //https 보안연결.

	//아래는 기존의 코드. net/http 기본 모듈 사용
	/*
		err := e.Start(":9090")
		if err != nil {
			e.Logger.Fatal(err)
		}
	*/

	/*
		fs := http.FileServer(http.Dir("./frontend/static"))
		http.Handle("/static/", http.StripPrefix("/static/", fs))

		http.HandleFunc("/", backend.Host_Index)
		http.HandleFunc("/check/", backend.Login_Check)
		http.HandleFunc("/main/", backend.User_Index)
		http.HandleFunc("/fail/", backend.Fail)

		//기존 코드에서 컨트롤러 패턴을 적용한 방식
		http.HandleFunc("/menu/", backend.Request_Handler)*/

	//기존 코드에서 컨트롤러 패턴을 적용하기 이전 방식
	/*
		http.HandleFunc("/notice_board/", backend.Request_Handler)
		http.HandleFunc("/notice_board_contents/", backend.Request_Handler)
		http.HandleFunc("/notice_board_write/", backend.Request_Handler)

		http.HandleFunc("/project/", backend.Request_Handler)
		http.HandleFunc("/project_contents/", backend.Request_Handler)
	*/

	/*
		log.Println("Listening on : 9090...")
		err := http.ListenAndServe(":9090", nil)

		if err != nil {
			log.Fatal("ListenAndServer : ", err)
		} else {
			fmt.Println("ListenAndServer Started! -> Port(9000)")
		}
	*/
}
