package backend

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

//handler.go 파일의 Request_handler에 의해 호출된 함수가 실행된다.

//EchoNoticeboardIndex is shows noticeboard tab index
/*
게시판의 첫 화면으로써 게시글 번호, 게시글 제목, 날짜, 조회 수가 출력된다.
페이징 기능이 적용되어 있다.
*/
func EchoNoticeboardIndex(c echo.Context) error {
	c.Request().ParseForm()
	respage := c.FormValue("Page")
	//rescount := c.FormValue("Count")

	//var n Interf = notice_board_view{}

	if respage != "" {
		/*
			intrespage, err := strconv.Atoi(respage)
			intrescount, err := strconv.Atoi(rescount)
			if err != nil {
				log.Fatal(err)
			}
			minint := (intrescount * 10) - (intrespage * 10)
			maxint := minint + 10
			minstr := strconv.Itoa(minint)
			maxstr := strconv.Itoa(maxint)
			var noticeviewstring = "SELECT *, (SELECT MAX(No) FROM notice_board_view) FROM notice_board_view WHERE No <=" + maxstr + " AND No >" + minstr + " ORDER BY No DESC limit 10;"
		*/

		intrespage, err := strconv.Atoi(respage)
		if err != nil {
			log.Fatal(err)
		}
		intrespage = (intrespage * 10) - 10
		page := strconv.Itoa(intrespage)
		var noticeviewstring = "SELECT *, (SELECT MAX(No) FROM notice_board_view) FROM notice_board_view WHERE No > 0 ORDER BY No DESC limit 10 OFFSET " + page + ";"

		result := NoticeSelectQuery(db1, noticeviewstring, "index")

		hostcookie, _ := c.Cookie("KKH")
		if hostcookie != nil {
			result = append(result, noticeboardcontentview{Cookie: "TRUE"})
		} else {
			result = append(result, noticeboardcontentview{Cookie: "FALSE"})
		}
		return c.Render(http.StatusOK, "notice_board.html", result)
	} else {
		var noticeviewstring = "SELECT *, (SELECT MAX(No) FROM notice_board_view) FROM notice_board_view WHERE No>(SELECT TRUNCATE((SELECT MAX(no)-1 FROM notice_board_view), -1) FROM dual) ORDER BY No DESC limit 10;"
		result := NoticeSelectQuery(db1, noticeviewstring, "index")

		hostcookie, _ := c.Cookie("KKH")
		fmt.Println(hostcookie)
		if hostcookie != nil {
			//result.Cookie = "TRUE"
			result = append(result, noticeboardcontentview{Cookie: "TRUE"})
		} else {
			//result.Cookie = "FALSE"
			result = append(result, noticeboardcontentview{Cookie: "FALSE"})
		}

		return c.Render(http.StatusOK, "notice_board.html", result)
	}
}

//EchoNoticeboardContentView is selected noticeboard content view and count the hits
/*
게시글을 선택하였을 때 보여지는 화면으로써 제목과 내용이 출력된다.
Utterances 댓글 기능과 다음 게시글, 이전 게시글로 바로 넘어갈 수 있는 기능, 그리고 24시간 쿠키를 적용하여 쿠키가 있을 경우에는 조회 수가 증가하지 않는 기능이 적용되어 있다.
*/
func EchoNoticeboardContentView(c echo.Context) error {
	c.Request().ParseForm()
	resno := c.Request().FormValue("No")
	if resno != "" {
		var noticeviewstring = ""
		var no = "SELECT MIN(No), MAX(No) FROM notice_board_view"
		minno, maxno := MaxSelectQuery(db1, no)
		minno = minno + 3
		convminno := strconv.Itoa(minno)
		convmaxno := strconv.Itoa(maxno)

		if (resno == convminno) && (resno == convmaxno) {
			noticeviewstring = "SELECT *, (SELECT MAX(No) FROM notice_board_view), (SELECT MIN(No) FROM notice_board_view), -2, (SELECT Title FROM notice_board_view WHERE No=-2), -1, (SELECT Title FROM notice_board_view WHERE No=-1) FROM notice_board_view WHERE No=" + resno + ";"
		} else if resno == convminno {
			noticeviewstring = "SELECT *, (SELECT MAX(No) FROM notice_board_view), (SELECT MIN(No) FROM notice_board_view), -2, (SELECT Title FROM notice_board_view WHERE No=-2), No+1, (SELECT Title FROM notice_board_view WHERE No=" + resno + "+1) FROM notice_board_view WHERE No=" + resno + ";"
		} else if resno == convmaxno {
			noticeviewstring = "SELECT *, (SELECT MAX(No) FROM notice_board_view), (SELECT MIN(No) FROM notice_board_view), No-1, (SELECT Title FROM notice_board_view WHERE No=" + resno + "-1), -1, (SELECT Title FROM notice_board_view WHERE No=-1) FROM notice_board_view WHERE No=" + resno + ";"
		} else {
			noticeviewstring = "SELECT *, (SELECT MAX(No) FROM notice_board_view), (SELECT MIN(No) FROM notice_board_view), No-1, (SELECT Title FROM notice_board_view WHERE No=" + resno + "-1), No+1, (SELECT Title FROM notice_board_view WHERE No=" + resno + "+1) FROM notice_board_view WHERE No=" + resno + ";"
		}

		result := NoticeSelectQuery(db1, noticeviewstring, "content")
		var noticecountupdate = "UPDATE notice_board_view SET Click=Click+1 WHERE No=" + resno + ";"

		cookie := readCookie(c, resno)

		if cookie == "cookie error" {
			writeCookie(c, resno, "no"+resno) //only used on 433 port(https) - chrome security policy.
			UpdateQuery(db1, noticecountupdate)
		}
		return c.Render(http.StatusOK, "notice_board_contents.html", result)
	}
	return c.Render(http.StatusOK, "error.html", 0)
}

//EchoNoticeboardWriteView is shows noticeboard write view and insert the written content into the database
/*
글쓰기 버튼을 눌렀을 때 보여지는 화면으로써 작성한 제목과 내용을 POST방식으로 DB에 저장한다.
내용 작성 창은 네이버의 스마트에디터2 를 적용하였다.
*/
func EchoNoticeboardWriteView(c echo.Context) error {
	if c.Request().Method == "POST" {
		dt := time.Now()
		fdt := dt.Format("2006-01-02")

		restitle := c.FormValue("title")
		rescontent := c.FormValue("ir1")
		fmt.Println(rescontent)
		var insertstring = "INSERT INTO notice_board_view (Title, Writer, Content, Date, Click) VALUES (" + "'" + restitle + "'" + ", '김광호', " + "'" + rescontent + "'" + ", " + "'" + fdt + "'" + ", " + "0" + ");"
		InsertQuery(db1, insertstring)
		http.Redirect(c.Response(), c.Request(), "/menu/?Handler=n_main", http.StatusFound)
	} else {
		return c.Render(http.StatusOK, "notice_board_write.html", "0")
	}
	return c.Render(http.StatusOK, "error.html", 0)
}
