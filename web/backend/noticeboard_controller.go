package backend

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

//handler.go 파일의 Request_handler에 의해 호출된 함수가 실행된다.

/*
게시판의 첫 화면으로써 게시글 번호, 게시글 제목, 날짜, 조회 수가 출력된다.
페이징 기능이 적용되어 있다.
*/
func Echo_Noticeboard_Index(c echo.Context) error {
	c.Request().ParseForm()
	respage := c.FormValue("Page")
	rescount := c.FormValue("Count")

	//var n Interf = notice_board_view{}

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
		var notice_view_string = "SELECT *, (SELECT MAX(No) FROM notice_board_view) FROM notice_board_view WHERE No <=" + max_str + " AND No >" + min_str + " ORDER BY No DESC limit 10;"

		result := SelectQuery(db1, notice_view_string)
		//result := SelectQuery(db1, notice_view_string)

		hostcookie, _ := c.Cookie("KKH")
		if hostcookie != nil {
			//result.Cookie = "TRUE"
			result = append(result, notice_board_view{Cookie: "TRUE"})
		} else {
			//result.Cookie = "FALSE"
			result = append(result, notice_board_view{Cookie: "FALSE"})
		}
		return c.Render(http.StatusOK, "notice_board.html", result)
		//return c.Render(http.StatusOK, "notice_board.html", result)
	} else {
		//var notice_view_string = "SELECT * FROM notice_board_view WHERE No<=(SELECT MAX(No) FROM notice_board_view) AND No>(SELECT TRUNCATE((SELECT MAX(no)-1 FROM notice_board_view), -1) FROM dual) ORDER BY No DESC limit 10;"
		var notice_view_string = "SELECT *, (SELECT MAX(No) FROM notice_board_view) FROM notice_board_view WHERE No>(SELECT TRUNCATE((SELECT MAX(no)-1 FROM notice_board_view), -1) FROM dual) ORDER BY No DESC limit 10;"
		result := SelectQuery(db1, notice_view_string)
		//result := SelectQuery(db1, notice_view_string)

		hostcookie, _ := c.Cookie("KKH")
		if hostcookie != nil {
			//result.Cookie = "TRUE"
			result = append(result, notice_board_view{Cookie: "TRUE"})
		} else {
			//result.Cookie = "FALSE"
			result = append(result, notice_board_view{Cookie: "FALSE"})
		}

		return c.Render(http.StatusOK, "notice_board.html", result)
	}
	return c.HTML(0, "ERROR")
}

/*
게시글을 선택하였을 때 보여지는 화면으로써 제목과 내용이 출력된다.
Utterances 댓글 기능과 다음 게시글, 이전 게시글로 바로 넘어갈 수 있는 기능, 그리고 24시간 쿠키를 적용하여 쿠키가 있을 경우에는 조회 수가 증가하지 않는 기능이 적용되어 있다.
*/

func Echo_Noticeboard_Content_View(c echo.Context) error {
	c.Request().ParseForm()
	resno := c.Request().FormValue("No")
	if resno != "" {
		var notice_view_string = ""
		var no = "SELECT MIN(No), MAX(No) FROM notice_board_view"
		minno, maxno := MaxSelectQuery(db1, no)
		minno = minno + 3
		convminno := strconv.Itoa(minno)
		convmaxno := strconv.Itoa(maxno)

		if (resno == convminno) && (resno == convmaxno) {
			notice_view_string = "SELECT *, (SELECT MAX(No) FROM notice_board_view), (SELECT MIN(No) FROM notice_board_view), -2, (SELECT Title FROM notice_board_view WHERE No=-2), -1, (SELECT Title FROM notice_board_view WHERE No=-1) FROM notice_board_view WHERE No=" + resno + ";"
		} else if resno == convminno {
			notice_view_string = "SELECT *, (SELECT MAX(No) FROM notice_board_view), (SELECT MIN(No) FROM notice_board_view), -2, (SELECT Title FROM notice_board_view WHERE No=-2), No+1, (SELECT Title FROM notice_board_view WHERE No=" + resno + "+1) FROM notice_board_view WHERE No=" + resno + ";"
		} else if resno == convmaxno {
			notice_view_string = "SELECT *, (SELECT MAX(No) FROM notice_board_view), (SELECT MIN(No) FROM notice_board_view), No-1, (SELECT Title FROM notice_board_view WHERE No=" + resno + "-1), -1, (SELECT Title FROM notice_board_view WHERE No=-1) FROM notice_board_view WHERE No=" + resno + ";"
		} else {
			notice_view_string = "SELECT *, (SELECT MAX(No) FROM notice_board_view), (SELECT MIN(No) FROM notice_board_view), No-1, (SELECT Title FROM notice_board_view WHERE No=" + resno + "-1), No+1, (SELECT Title FROM notice_board_view WHERE No=" + resno + "+1) FROM notice_board_view WHERE No=" + resno + ";"
		}
		//var notice_view_string = "SELECT *, (SELECT MAX(No) FROM notice_board_view) FROM notice_board_view WHERE No=" + resno + ";"
		//result := SelectQuery(db1, notice_view_string)

		result := Noticeboard_ContentQuery(db1, notice_view_string)
		var notice_count_update = "UPDATE notice_board_view SET Click=Click+1 WHERE No=" + resno + ";"

		cookie, _ := c.Cookie(resno)
		if cookie == nil {
			writeCookie(c, resno, "no"+resno)
			UpdateQuery(db1, notice_count_update)
		}
		//cookie.Value = cookie.Value + "test"
		//fmt.Println("cookie test : ", cookie.Value)
		//c.SetCookie(cookie)
		//UpdateQuery(db1, notice_count_update)
		//Echo_Noticeboard_Content_View(c)
		return c.Render(http.StatusOK, "notice_board_contents.html", result)
	}
	return c.HTML(0, "ERROR")
}

/*
글쓰기 버튼을 눌렀을 때 보여지는 화면으로써 작성한 제목과 내용을 POST방식으로 DB에 저장한다.
내용 작성 창은 네이버의 스마트에디터2 를 적용하였다.
*/
func Echo_Noticeboard_Write_View(c echo.Context) error {
	if c.Request().Method == "POST" {
		dt := time.Now()
		f_dt := dt.Format("2006-01-02")

		restitle := c.FormValue("title")
		rescontent := c.FormValue("ir1")
		fmt.Println(rescontent)
		var insert_string = "INSERT INTO notice_board_view (Title, Writer, Content, Date, Click) VALUES (" + "'" + restitle + "'" + ", '김광호', " + "'" + rescontent + "'" + ", " + "'" + f_dt + "'" + ", " + "0" + ");"
		InsertQuery(db1, insert_string)
		http.Redirect(c.Response(), c.Request(), "/menu/?Handler=n_main", http.StatusFound)
	} else {
		return c.Render(http.StatusOK, "notice_board_write.html", "0")
	}
	return c.String(0, "ERROR")
}

// 이전 방식의 게시판 첫 화면 및 페이징 기능
func Noticeboard_Index(w http.ResponseWriter, r *http.Request) {
	noticeboardTemplate, _ := template.ParseFiles("frontend/notice_board.html", header, footer, leftside)
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
		var notice_view_string = "SELECT *, (SELECT MAX(No) FROM notice_board_view) FROM notice_board_view WHERE No <=" + max_str + " AND No >" + min_str + " ORDER BY No DESC limit 10;"
		result := SelectQuery(db1, notice_view_string)
		noticeboardTemplate.Execute(w, result)
	} else {
		//var notice_view_string = "SELECT * FROM notice_board_view WHERE No<=(SELECT MAX(No) FROM notice_board_view) AND No>(SELECT TRUNCATE((SELECT MAX(no)-1 FROM notice_board_view), -1) FROM dual) ORDER BY No DESC limit 10;"
		var notice_view_string = "SELECT *, (SELECT MAX(No) FROM notice_board_view) FROM notice_board_view WHERE No>(SELECT TRUNCATE((SELECT MAX(no)-1 FROM notice_board_view), -1) FROM dual) ORDER BY No DESC limit 10;"
		result := SelectQuery(db1, notice_view_string)
		noticeboardTemplate.Execute(w, result)
	}
}

// 글을 클릭했을 때 보여지는 내용 및 조회 수 계산
func Noticeboard_Content_View(w http.ResponseWriter, r *http.Request) {
	noticeboardcontentsTemplate, _ := template.ParseFiles("frontend/notice_board_contents.html", header, footer, leftside)
	r.ParseForm()
	resno := r.FormValue("No")
	if resno != "" {
		var notice_view_string = "SELECT *, (SELECT MAX(No) FROM notice_board_view) FROM notice_board_view WHERE No=" + resno + ";"
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
	noticeboardwriteTemplate, _ := template.ParseFiles("frontend/notice_board_write.html", header, footer, leftside)

	r.ParseForm()
	if r.Method == "POST" {
		dt := time.Now()
		f_dt := dt.Format("2006-01-02")

		restitle := r.FormValue("title")
		rescontent := r.FormValue("ir1")
		fmt.Println(rescontent)
		var insert_string = "INSERT INTO notice_board_view (Title, Writer, Content, Date, Click) VALUES (" + "'" + restitle + "'" + ", '김광호', " + "'" + rescontent + "'" + ", " + "'" + f_dt + "'" + ", " + "0" + ");"
		InsertQuery(db1, insert_string)
		http.Redirect(w, r, "/menu/?Handler=n_main", http.StatusFound)
	} else {
		noticeboardwriteTemplate.Execute(w, nil)
	}
}
