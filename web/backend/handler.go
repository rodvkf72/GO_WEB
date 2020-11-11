package backend

import (
	"net/http"

	"github.com/labstack/echo"
)

/*
GET방식으로 넘겨지는 Hnadler의 값을 가지고 global_info.go 파일에 있는 static 함수를 호출하여 리턴된 값을 가지고 해당하는 함수를 호출
*/
func Echo_Request_Handler(c echo.Context) error {
	reshandler := c.FormValue("Handler")
	s_handle := static(reshandler)

	switch s_handle {
	case "1":
		return Echo_Noticeboard_Index(c)
	case "2":
		return Echo_Noticeboard_Content_View(c)
	case "3":
		return Echo_Noticeboard_Write_View(c)

	case "11":
		return Echo_Project_Index(c)
	case "12":
		return Echo_Project_Content_View(c)
	case "13":
		return Echo_Project_Write_View(c)

	case "21":
		return Echo_Game_Index(c)
	case "23":
		return Echo_Game_Write_View(c)

	case "31":
		return Echo_Baekjoon_Index(c)
	case "32":
		return Echo_Baekjoon_Content_View(c)
	case "33":
		return Echo_Baekjoon_Write_View(c)
	case "34":
		return Echo_Baekjoon_Search(c)
	}
	return c.String(0, "ERROR")
}

// 이전 방식의 메뉴 버튼 클릭 시 화면 지정
func Request_Handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	reshandler := r.FormValue("Handler")
	s_handle := static(reshandler)
	switch s_handle {
	case "1":
		Noticeboard_Index(w, r)
	case "2":
		Noticeboard_Content_View(w, r)
	case "3":
		Noticeboard_Write_View(w, r)

	case "11":
		Project_Index(w, r)
	case "12":
		Project_Content_View(w, r)
	case "13":
		Project_Write_View(w, r)

	case "21":
		Game_Index(w, r)
	/*
		case "22":
			Game_Content_View(w, r)
	*/
	case "23":
		Game_Write_View(w, r)
	}
}
