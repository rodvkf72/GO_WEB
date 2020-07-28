package backend

import (
	"net/http"
)

// 메뉴 버튼 클릭 시 화면을 지정해줌
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