package backend

import (
	"database/sql"
	"log"
)

// DB update 문
func UpdateQuery(db dbInfo, query string) {
	dataSource := db.user + ":" + db.pwd + "@tcp(" + db.url + ")/" + db.database
	conn, err := sql.Open(db.engine, dataSource)
	if err != nil {
		log.Fatal(err)
	}
	conn.Exec(query)
	conn.Close()
}