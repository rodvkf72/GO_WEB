package backend

import (
	"net/http"
	"html/template"
)

func Game_Index(w http.ResponseWriter, r *http.Request) {
	gameTemplate, _ := template.ParseFiles("frontend/game.html", "frontend/header.html", "frontend/footer.html")

	var game_view_string = "SELECT * FROM game_view"
	result := GameSelectQuery(db1, game_view_string)
	gameTemplate.Execute(w, result)
}

func Game_Write_View(w http.ResponseWriter, r *http.Request) {
	gamewriteTemplate, _ := template.ParseFiles("frontend/game_write.html", "frontend/header.html", "frontend/footer.html")

	r.ParseForm()
	if r.Method == "POST" {
		resgame := r.FormValue("game")
		resroot := r.FormValue("root")
		var insert_string = "INSERT INTO game_view (Game, Root) VALUES (" + "'" + resgame + "'" + "," + "'" + resroot + "'" + ");"
		InsertQuery(db1, insert_string)
		http.Redirect(w, r, "/menu/?Handler=g_main", http.StatusFound)
	} else {
		gamewriteTemplate.Execute(w, nil)
	}
}