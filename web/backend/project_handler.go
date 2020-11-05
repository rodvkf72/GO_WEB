package backend

import (
	"net/http"
)

//Echo Framework 에선 쓰이지 않음
func Project_Handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	reshandler := r.FormValue("Handler")
	s_handle := static(reshandler)
	switch s_handle {
	case "11":
		Project_Index(w, r)
	case "12":
		Project_Content_View(w, r)
	case "13":
		Project_Write_View(w, r)
	}
}