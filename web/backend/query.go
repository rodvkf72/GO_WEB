package backend

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

//DB query 문
//값을 구조체로 전달하다 보니 Scan에서 갯수가 일치하지 않는 경우 에러가 출력됨. 따라서 구조체가 다른 경우 각각의 함수를 구현

//InsertQuery is mysql Insert Query
func InsertQuery(db dbInfo, query string) {
	dataSource := db.user + ":" + db.pwd + "@tcp(" + db.url + ")/" + db.database
	conn, err := sql.Open(db.engine, dataSource)
	if err != nil {
		log.Fatal(err)
	}
	result, err := conn.Exec(query)

	if err != nil {
		log.Fatal(err)
	}
	nRow, err := result.RowsAffected()
	fmt.Println("insert count : ", nRow)
	conn.Close()
}

//UpdateQuery is mysql Update Query
func UpdateQuery(db dbInfo, query string) {
	dataSource := db.user + ":" + db.pwd + "@tcp(" + db.url + ")/" + db.database
	conn, err := sql.Open(db.engine, dataSource)
	if err != nil {
		log.Fatal(err)
	}
	conn.Exec(query)
	conn.Close()
}

//NoticeSelectQuery is notice board index view and notice board content view query
/*
게시글의 첫 화면, 선택된 게시글 내용을 출력하기 위한 select 문
다음 글, 이전 글 기능을 위해 구조체에 추가된 변수가 많다.
*/
func NoticeSelectQuery(db dbInfo, query string, sel string) []notice_board_content_view {
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
		if sel == "index" {
			err := rows.Scan(&no, &title, &writer, &content, &date, &click, &maxno)
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
			Notice_board_content_views = append(Notice_board_content_views, Notice_board_content_view)
		} else {
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
	}
	defer conn.Close()
	return Notice_board_content_views
}

//MaxSelectQuery is find the ninimum and maximum value of the post number
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

//GameSelectQuery is game tab inex view query
/*
게임 탭의 첫 화면 출력을 위한 select 문
*/
func GameSelectQuery(db dbInfo, query string, sel string) []game_content_view {
	dataSource := db.user + ":" + db.pwd + "@tcp(" + db.url + ")/" + db.database
	conn, err := sql.Open(db.engine, dataSource)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := conn.Query(query)

	Game_view := game_content_view{}
	Game_views := []game_content_view{}

	for rows.Next() {
		var no int
		var game, types, root, content string
		if sel == "index" {
			err := rows.Scan(&no, &game, &types, &root)
			if err != nil {
				log.Fatal(err)
			}
			Game_view.No = no
			Game_view.Game = game
			Game_view.Type = types
			Game_view.Root = root
			Game_views = append(Game_views, Game_view)
		} else {
			err := rows.Scan(&no, &game, &types, &root, &content)
			if err != nil {
				log.Fatal(err)
			}
			Game_view.No = no
			Game_view.Game = game
			Game_view.Type = types
			Game_view.Root = root
			Game_view.Content = content
			Game_views = append(Game_views, Game_view)
		}
	}
	defer conn.Close()
	return Game_views
}

//GameContentQuery is game tab content view query
func GameContentQuery(db dbInfo, query string) []game_content_view {
	dataSource := db.user + ":" + db.pwd + "@tcp(" + db.url + ")/" + db.database
	conn, err := sql.Open(db.engine, dataSource)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := conn.Query(query)

	Game_view := game_content_view{}
	Game_views := []game_content_view{}

	for rows.Next() {
		var no int
		var game, types, root, content string
		err := rows.Scan(&no, &game, &types, &root, &content)
		if err != nil {
			log.Fatal(err)
		}
		Game_view.No = no
		Game_view.Game = game
		Game_view.Type = types
		Game_view.Root = root
		Game_view.Content = content
		Game_views = append(Game_views, Game_view)
	}
	defer conn.Close()
	return Game_views
}

//BaekjoonSelectQuery is coding tab index view query
func BaekjoonSelectQuery(db dbInfo, query string, sel string) []baekjoon_content_view {
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
		if sel == "index" {
			err := rows.Scan(&no, &title)
			if err != nil {
				log.Fatal(err)
			}
			Baekjoon_content_view.No = no
			Baekjoon_content_view.Title = title
			Baekjoon_content_views = append(Baekjoon_content_views, Baekjoon_content_view)
		} else {
			err := rows.Scan(&no, &title, &content)
			if err != nil {
				log.Fatal(err)
			}
			Baekjoon_content_view.No = no
			Baekjoon_content_view.Title = title
			Baekjoon_content_view.Content = content
			Baekjoon_content_views = append(Baekjoon_content_views, Baekjoon_content_view)
		}
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
