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
func NoticeSelectQuery(db dbInfo, query string, sel string) []noticeboardcontentview {
	dataSource := db.user + ":" + db.pwd + "@tcp(" + db.url + ")/" + db.database
	conn, err := sql.Open(db.engine, dataSource)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := conn.Query(query)
	Noticeboardcontentview := noticeboardcontentview{}
	Noticeboardcontentviews := []noticeboardcontentview{}
	for rows.Next() {
		var no, click, maxno, minno, minusno, plusno int
		var title, writer, content, date, minustitle, plustitle string
		if sel == "index" {
			err := rows.Scan(&no, &title, &writer, &content, &date, &click, &maxno)
			if err != nil {
				log.Fatal(err)
			}
			Noticeboardcontentview.No = no
			Noticeboardcontentview.Title = title
			Noticeboardcontentview.Writer = writer
			Noticeboardcontentview.Content = strings.Replace(content, "\r\n", "\n", 1)
			Noticeboardcontentview.Date = date
			Noticeboardcontentview.Click = click
			Noticeboardcontentview.Maxno = maxno
			Noticeboardcontentviews = append(Noticeboardcontentviews, Noticeboardcontentview)
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
			Noticeboardcontentview.No = no
			Noticeboardcontentview.Title = title
			Noticeboardcontentview.Writer = writer
			Noticeboardcontentview.Content = strings.Replace(content, "\r\n", "\n", 1)
			Noticeboardcontentview.Date = date
			Noticeboardcontentview.Click = click
			Noticeboardcontentview.Maxno = maxno
			Noticeboardcontentview.Minno = minno + 3
			Noticeboardcontentview.Minusno = minusno
			Noticeboardcontentview.Minustitle = minustitle
			Noticeboardcontentview.Plusno = plusno
			Noticeboardcontentview.Plustitle = plustitle
			Noticeboardcontentviews = append(Noticeboardcontentviews, Noticeboardcontentview)
		}
	}
	defer conn.Close()
	return Noticeboardcontentviews
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
func GameSelectQuery(db dbInfo, query string, sel string) []gamecontentview {
	dataSource := db.user + ":" + db.pwd + "@tcp(" + db.url + ")/" + db.database
	conn, err := sql.Open(db.engine, dataSource)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := conn.Query(query)

	Gameview := gamecontentview{}
	Gameviews := []gamecontentview{}

	for rows.Next() {
		var no int
		var game, types, root, content string
		if sel == "index" {
			err := rows.Scan(&no, &game, &types, &root)
			if err != nil {
				log.Fatal(err)
			}
			Gameview.No = no
			Gameview.Game = game
			Gameview.Type = types
			Gameview.Root = root
			Gameviews = append(Gameviews, Gameview)
		} else {
			err := rows.Scan(&no, &game, &types, &root, &content)
			if err != nil {
				log.Fatal(err)
			}
			Gameview.No = no
			Gameview.Game = game
			Gameview.Type = types
			Gameview.Root = root
			Gameview.Content = content
			Gameviews = append(Gameviews, Gameview)
		}
	}
	defer conn.Close()
	return Gameviews
}

//GameContentQuery is game tab content view query
func GameContentQuery(db dbInfo, query string) []gamecontentview {
	dataSource := db.user + ":" + db.pwd + "@tcp(" + db.url + ")/" + db.database
	conn, err := sql.Open(db.engine, dataSource)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := conn.Query(query)

	Gameview := gamecontentview{}
	Gameviews := []gamecontentview{}

	for rows.Next() {
		var no int
		var game, types, root, content string
		err := rows.Scan(&no, &game, &types, &root, &content)
		if err != nil {
			log.Fatal(err)
		}
		Gameview.No = no
		Gameview.Game = game
		Gameview.Type = types
		Gameview.Root = root
		Gameview.Content = content
		Gameviews = append(Gameviews, Gameview)
	}
	defer conn.Close()
	return Gameviews
}

//BaekjoonSelectQuery is coding tab index view query
func BaekjoonSelectQuery(db dbInfo, query string, sel string) []baekjooncontentview {
	dataSource := db.user + ":" + db.pwd + "@tcp(" + db.url + ")/" + db.database
	conn, err := sql.Open(db.engine, dataSource)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := conn.Query(query)

	Baekjooncontentview := baekjooncontentview{}
	Baekjooncontentviews := []baekjooncontentview{}

	for rows.Next() {
		var no, count int
		var title, content string
		if sel == "index" {
			err := rows.Scan(&no, &title, &count)
			if err != nil {
				log.Fatal(err)
			}
			Baekjooncontentview.No = no
			Baekjooncontentview.Title = title
			Baekjooncontentview.Count = count
			Baekjooncontentviews = append(Baekjooncontentviews, Baekjooncontentview)
		} else {
			err := rows.Scan(&no, &title, &content)
			if err != nil {
				log.Fatal(err)
			}
			Baekjooncontentview.No = no
			Baekjooncontentview.Title = title
			Baekjooncontentview.Content = content
			Baekjooncontentviews = append(Baekjooncontentviews, Baekjooncontentview)
		}
	}
	defer conn.Close()
	return Baekjooncontentviews
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
