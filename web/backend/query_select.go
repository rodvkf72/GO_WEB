package backend

import (
	"database/sql"
	"log"
	"strings"
)

//DB select 문
//값을 구조체로 전달하다 보니 Scan에서 갯수가 일치하지 않는 경우 에러가 출력됨. 따라서 구조체가 다른 경우 각각의 함수를 구현

/*
func (n notice_board_view) SelectQuery(db dbInfo, query string) []notice_board_view {
	dataSource := db.user + ":" + db.pwd + "@tcp(" + db.url + ")/" + db.database
	conn, err := sql.Open(db.engine, dataSource)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := conn.Query(query)

	//Notice_board_view := notice_board_view{}
	Notice_board_views := []notice_board_view{}

	for rows.Next() {
		var no, click, maxno int
		var title, writer, content, date string
		err := rows.Scan(&no, &title, &writer, &content, &date, &click, &maxno)
		if err != nil {
			log.Fatal(err)
		}
		n.No = no
		n.Title = title
		n.Writer = writer
		n.Content = strings.Replace(content, "\r\n", "\n", 1)
		n.Date = date
		n.Click = click
		n.Maxno = maxno
		Notice_board_views = append(Notice_board_views, n)
	}
	defer conn.Close()
	return Notice_board_views
}
*/

/*
게시글의 첫 화면을 출력하기 위한 select 문
*/
func SelectQuery(db dbInfo, query string) []notice_board_view {
	dataSource := db.user + ":" + db.pwd + "@tcp(" + db.url + ")/" + db.database
	conn, err := sql.Open(db.engine, dataSource)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := conn.Query(query)

	Notice_board_view := notice_board_view{}
	Notice_board_views := []notice_board_view{}

	for rows.Next() {
		var no, click, maxno int
		var title, writer, content, date string
		err := rows.Scan(&no, &title, &writer, &content, &date, &click, &maxno)
		if err != nil {
			log.Fatal(err)
		}
		Notice_board_view.No = no
		Notice_board_view.Title = title
		Notice_board_view.Writer = writer
		Notice_board_view.Content = strings.Replace(content, "\r\n", "\n", 1)
		Notice_board_view.Date = date
		Notice_board_view.Click = click
		Notice_board_view.Maxno = maxno
		Notice_board_views = append(Notice_board_views, Notice_board_view)
	}
	defer conn.Close()
	return Notice_board_views
}

/*
게시글의 페이징 기능을 위한 최소값과 최대값을 가져오는 select 문
*/
func MaxSelectQuery(db dbInfo, query string) (int, int) {
	dataSource := db.user + ":" + db.pwd + "@tcp(" + db.url + ")/" + db.database
	conn, err := sql.Open(db.engine, dataSource)
	var minno, maxno int
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	err = conn.QueryRow(query).Scan(&minno, &maxno)
	if err != nil {
		log.Fatal(err)
	}
	return minno, maxno
}

/*
선택된 게시글의 내용을 출력하기 위한 select 문
다음 글, 이전 글 기능을 위해 구조체에 추가된 변수가 많다.
*/
func Noticeboard_ContentQuery(db dbInfo, query string) []notice_board_content_view {
	dataSource := db.user + ":" + db.pwd + "@tcp(" + db.url + ")/" + db.database
	conn, err := sql.Open(db.engine, dataSource)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := conn.Query(query)
	Notice_board_content_view := notice_board_content_view{}
	Notice_board_content_views := []notice_board_content_view{}
	for rows.Next() {
		var no, click, maxno, minno, minusno, plusno int
		var title, writer, content, date, minustitle, plustitle string
		err := rows.Scan(&no, &title, &writer, &content, &date, &click, &maxno, &minno, &minusno, &minustitle, &plusno, &plustitle)
		if plusno == 0 {
			plusno = 0
		}
		if plustitle == "" {
			plustitle = "다음 글이 없습니다."
		}
		if err != nil {
			log.Fatal(err)
		}
		Notice_board_content_view.No = no
		Notice_board_content_view.Title = title
		Notice_board_content_view.Writer = writer
		Notice_board_content_view.Content = strings.Replace(content, "\r\n", "\n", 1)
		Notice_board_content_view.Date = date
		Notice_board_content_view.Click = click
		Notice_board_content_view.Maxno = maxno
		Notice_board_content_view.Minno = minno + 3
		Notice_board_content_view.Minusno = minusno
		Notice_board_content_view.Minustitle = minustitle
		Notice_board_content_view.Plusno = plusno
		Notice_board_content_view.Plustitle = plustitle
		Notice_board_content_views = append(Notice_board_content_views, Notice_board_content_view)
	}
	defer conn.Close()
	return Notice_board_content_views
}

/*
게임 탭의 첫 화면 출력을 위한 select 문
*/
func GameSelectQuery(db dbInfo, query string) []game_view {
	dataSource := db.user + ":" + db.pwd + "@tcp(" + db.url + ")/" + db.database
	conn, err := sql.Open(db.engine, dataSource)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := conn.Query(query)

	Game_view := game_view{}
	Game_views := []game_view{}

	for rows.Next() {
		var no int
		var name, types, root string
		err := rows.Scan(&no, &name, &types, &root)
		if err != nil {
			log.Fatal(err)
		}
		Game_view.No = no
		Game_view.Name = name
		Game_view.Type = types
		Game_view.Root = root
		Game_views = append(Game_views, Game_view)
	}
	defer conn.Close()
	return Game_views
}

func BaekjoonSelectQuery(db dbInfo, query string) []baekjoon_view {
	dataSource := db.user + ":" + db.pwd + "@tcp(" + db.url + ")/" + db.database
	conn, err := sql.Open(db.engine, dataSource)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := conn.Query(query)

	Baekjoon_view := baekjoon_view{}
	Baekjoon_views := []baekjoon_view{}

	for rows.Next() {
		var no int
		var title string
		err := rows.Scan(&no, &title)
		if err != nil {
			log.Fatal(err)
		}
		Baekjoon_view.No = no
		Baekjoon_view.Title = title
		Baekjoon_views = append(Baekjoon_views, Baekjoon_view)
	}
	defer conn.Close()
	return Baekjoon_views
}

func BaekjoonContentQuery(db dbInfo, query string) []baekjoon_content_view {
	dataSource := db.user + ":" + db.pwd + "@tcp(" + db.url + ")/" + db.database
	conn, err := sql.Open(db.engine, dataSource)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := conn.Query(query)

	Baekjoon_content_view := baekjoon_content_view{}
	Baekjoon_content_views := []baekjoon_content_view{}

	for rows.Next() {
		var no int
		var title, content string
		err := rows.Scan(&no, &title, &content)
		if err != nil {
			log.Fatal(err)
		}
		Baekjoon_content_view.No = no
		Baekjoon_content_view.Title = title
		Baekjoon_content_view.Content = content
		Baekjoon_content_views = append(Baekjoon_content_views, Baekjoon_content_view)
	}
	defer conn.Close()
	return Baekjoon_content_views
}

/* 단수 쿼리의 경우
dataSource := db.user + ":" + db.pwd + "@tcp(" + db.url + ")/" + db.database
conn, err := sql.Open(db.engine, dataSource)
if err != nil {
	log.Fatal(err)
}
defer conn.Close()
err = conn.QueryRow(query).Scan(&name)
if err != nil {
	log.Fatal(err)
}
fmt.Println(name)
return name
*/
