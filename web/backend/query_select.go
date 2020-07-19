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