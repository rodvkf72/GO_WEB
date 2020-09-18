package backend

import (
	"github.com/labstack/echo"
	"html/template"
	"net/http"
)

func Echo_Game_Index(c echo.Context) error {
	var game_view_string = "SELECT * FROM game_view"
	result := GameSelectQuery(db1, game_view_string)
	return c.Render(http.StatusOK, "game.html", result)
}

func Echo_Game_Write_View(c echo.Context) error {
	if c.Request().Method == "POST" {
		resgame := c.FormValue("game")
		restype := c.FormValue("type")
		resroot := c.FormValue("root")
		var insert_string = "INSERT INTO game_view (Game, Type, Root) VALUES (" + "'" + resgame + "'" + "," + "'" + restype + "'" + "," + "'" + resroot + "'" + ");"
		InsertQuery(db1, insert_string)
		http.Redirect(c.Response(), c.Request(), "/menu/?Handler=g_main", http.StatusFound)
	} else {
		return c.Render(http.StatusOK, "game_write.html", "0")
	}
	return c.String(0, "ERROR")
}

func Game_Index(w http.ResponseWriter, r *http.Request) {
	gameTemplate, _ := template.ParseFiles("frontend/game.html", header, footer, leftside)

	var game_view_string = "SELECT * FROM game_view"
	result := GameSelectQuery(db1, game_view_string)
	gameTemplate.Execute(w, result)
}

func Game_Write_View(w http.ResponseWriter, r *http.Request) {
	gamewriteTemplate, _ := template.ParseFiles("frontend/game_write.html", header, footer, leftside)

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