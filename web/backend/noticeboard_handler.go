package backend

import (
	"net/http"
)

func Noticeboard_Handler(w http.ResponseWriter, r *http.Request) {
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
	}
}