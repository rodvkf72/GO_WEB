package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"html/template"
	"io"
	"net/http"
	"web/backend"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render (w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	t := &Template {
		templates: template.Must(template.ParseGlob("./frontend/*.html")),
	}
	fs := http.FileServer(http.Dir("./frontend/static"))

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/static/", "public")
	e.Renderer = t

	e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", fs)))
	e.POST("/static/*", echo.WrapHandler(http.StripPrefix("/static/", fs)))
	//e.POST("/static/smarteditor2/sample/photo_uploader/attach_photo.js", echo.WrapHandler(http.StripPrefix("/static/smartedit2/", fs)))
	e.GET("/*", backend.Echo_Host_Index)
	e.POST("/check/*", backend.Echo_Login_Check)
	e.GET("/main/*", backend.Echo_User_Index)
	e.GET("/fail/*", backend.Echo_Fail)
	e.GET("/menu/*", backend.Echo_Request_Handler)
	e.POST("/menu/g_write", backend.Echo_Game_Write_View)
	e.POST("/menu/n_write", backend.Echo_Noticeboard_Write_View)
	e.POST("/single_img_upload/", backend.SingleImgUpload)
	e.POST("/multi_img_upload/", backend.MultiImgUpload)


	/*e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "/static/main.html", echo.Map{"title" : "Page file title!!"})
	})*/
	err := e.Start(":9090")
	if err != nil {
		e.Logger.Fatal(err)
	}

	/*
	fs := http.FileServer(http.Dir("./frontend/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", backend.Host_Index)
	http.HandleFunc("/check/", backend.Login_Check)
	http.HandleFunc("/main/", backend.User_Index)
	http.HandleFunc("/fail/", backend.Fail)

	http.HandleFunc("/menu/", backend.Request_Handler)*/

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
