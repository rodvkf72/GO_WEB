package backend

import (
	"html/template"
	"net/http"
)

func Game_Index(w http.ResponseWriter, r *http.Request) {
	gameTemplate, _ := template.ParseFiles("frontend/game.html", header, footer)

	var game_view_string = "SELECT * FROM game_view"
	result := GameSelectQuery(db1, game_view_string)
	gameTemplate.Execute(w, result)
}

func Game_Write_View(w http.ResponseWriter, r *http.Request) {
	gamewriteTemplate, _ := template.ParseFiles("frontend/game_write.html", header, footer)

	r.ParseForm()
	if r.Method == "POST" {
		resgame := r.FormValue("game")
		restype := r.FormValue("type")
		resroot := r.FormValue("root")
		var insert_string = "INSERT INTO game_view (Game, Type, Root) VALUES (" + "'" + resgame + "'" + "," + "'" + restype + "'" + "," + "'" + resroot + "'" + ");"
		InsertQuery(db1, insert_string)
		http.Redirect(w, r, "/menu/?Handler=g_main", http.StatusFound)
	} else {
		gamewriteTemplate.Execute(w, nil)
	}
}
