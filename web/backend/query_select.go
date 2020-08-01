package backend

import (
	"database/sql"
	"log"
	"strings"
)

// DB select 문
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
		var no, click int
		var title, writer, content, date string
		err := rows.Scan(&no, &title, &writer, &content, &date, &click)
		if err != nil {
			log.Fatal(err)
		}
		Notice_board_view.No = no
		Notice_board_view.Title = title
		Notice_board_view.Writer = writer
		Notice_board_view.Content = strings.Replace(content,"\r\n", "\n", 1)
		Notice_board_view.Date = date
		Notice_board_view.Click = click
		Notice_board_views = append(Notice_board_views, Notice_board_view)
	}
	defer conn.Close()
	return Notice_board_views
}

func NoSelectQuery(db dbInfo, query string) []notice_board_total_no {
	dataSource := db.user + ":" + db.pwd + "@tcp(" + db.url + ")/" + db.database
	conn, err := sql.Open(db.engine, dataSource)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := conn.Query(query)
	Notice_board_total_no := notice_board_total_no{}
	Notice_board_total_nos := []notice_board_total_no{}
	for rows.Next(){
		var total_no int
		err := rows.Scan(&total_no)
		if err != nil {
			log.Fatal(err)
		}
		Notice_board_total_no.Total_no = total_no
		Notice_board_total_nos = append(Notice_board_total_nos, Notice_board_total_no)
	}
	defer conn.Close()
	return Notice_board_total_nos
}

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
		var name, root string
		err := rows.Scan(&no, &name, &root)
		if err != nil {
			log.Fatal(err)
		}
		Game_view.No = no
		Game_view.Name = name
		Game_view.Root = root
		Game_views = append(Game_views, Game_view)
	}
	defer conn.Close()
	return Game_views
}

/* 단수 쿼리의 경우
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