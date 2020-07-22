package backend

import (
	"html/template"
	"net/http"
	"strconv"
	"log"
	"fmt"
	"time"
)

// 게시판 첫 화면 및 페이징 기능
func Noticeboard_Index(w http.ResponseWriter, r *http.Request) {
	noticeboardTemplate, _ := template.ParseFiles("frontend/notice_board.html", "frontend/header.html", "frontend/footer.html")
	r.ParseForm()
	respage := r.FormValue("Page")
	rescount := r.FormValue("Count")
	if respage != "" {
		int_respage, err := strconv.Atoi(respage)
		int_rescount, err := strconv.Atoi(rescount)
		if err != nil {
			log.Fatal(err)
		}
		min_int := (int_rescount * 10) - (int_respage * 10)
		max_int := min_int + 10
		min_str := strconv.Itoa(min_int)
		max_str := strconv.Itoa(max_int)
		fmt.Println(min_str)
		fmt.Println(max_str)
		var notice_view_string = "SELECT * FROM notice_board_view WHERE No <=" + max_str + " AND No >=" + min_str + " ORDER BY No DESC limit 10;"
		result := SelectQuery(db1, notice_view_string)
		noticeboardTemplate.Execute(w, result)
	} else {
		var notice_view_string = "SELECT * FROM notice_board_view ORDER BY No DESC limit 10;"
		result := SelectQuery(db1, notice_view_string)
		noticeboardTemplate.Execute(w, result)
	}
}

// 글을 클릭했을 때 보여지는 내용 및 조회 수 계산
func Noticeboard_Content_View(w http.ResponseWriter, r *http.Request) {
	noticeboardcontentsTemplate, _ := template.ParseFiles("frontend/notice_board_contents.html", "frontend/header.html", "frontend/footer.html")
	r.ParseForm()
	resno := r.FormValue("No")
	if resno != "" {
		var notice_view_string = "SELECT * FROM notice_board_view WHERE No=" + resno + ";"
		result := SelectQuery(db1, notice_view_string)
		var notice_count_update = "UPDATE notice_board_view SET Click=Click+1 WHERE No=" + resno + ";"
		UpdateQuery(db1, notice_count_update)
		noticeboardcontentsTemplate.Execute(w, result)
	}
	
	//html 요청 부분에서 name 값을 가지고 출력
	//resid := r.FormValue("No")
	//respw := r.FormValue("pw")
	//fmt.Println(resid)
	//fmt.Println(respw)

	//var notice_view_string = "SELECT * FROM notice_board_view WHERE No=" + s_resno + ";"
	//result := SelectQuery(db1, notice_view_string)
	//noticeboardcontentsTemplate.Execute(w, nil)
}

// 게시글 작성 시 날짜형식을 기호에 맞게 포맷하여 데이터베이스에 삽입
func Noticeboard_Write_View(w http.ResponseWriter, r *http.Request) {
	noticeboardwriteTemplate, _ := template.ParseFiles("frontend/notice_board_write.html", "frontend/header.html", "frontend/footer.html")

	r.ParseForm()
	if r.Method == "POST" {
		dt := time.Now()
		f_dt := dt.Format("2006-01-02")

		restitle := r.FormValue("title")
		rescontent := r.FormValue("content")
		var insert_string = "INSERT INTO notice_board_view (Title, Writer, Content, Date, Click) VALUES (" + "'" + restitle + "'" + ", '김광호', " + "'" + rescontent + "'" + ", " + "'" + f_dt + "'" + ", " + "0" + ");"
		InsertQuery(db1, insert_string)
		http.Redirect(w, r, "/notice_board/?Handler=n_main", http.StatusFound)
	} else {
		noticeboardwriteTemplate.Execute(w, nil)
	}
}