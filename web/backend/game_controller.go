package backend

import (
	"net/http"

	"github.com/labstack/echo"
)

//EchoGameIndex is shows the game tab index
func EchoGameIndex(c echo.Context) error {
	var gameviewstring = "SELECT no, game, type, root FROM game_view"
	result := GameSelectQuery(db1, gameviewstring, "index")

	hostcookie, _ := c.Cookie("KKH")
	if hostcookie != nil {
		result = append(result, gamecontentview{Cookie: "TRUE"})
	} else {
		result = append(result, gamecontentview{Cookie: "FALSE"})
	}

	return c.Render(http.StatusOK, "game.html", result)
}

//EchoGameContentView is shows selected game content
func EchoGameContentView(c echo.Context) error {
	resno := c.FormValue("no")
	if resno != "" {
		var gamecontentview = "SELECT * FROM game_view WHERE no=" + "'" + resno + "'" + ";"
		result := GameContentQuery(db1, gamecontentview)
		return c.Render(http.StatusOK, "game_contents.html", result)
	} else {
		return c.Render(http.StatusOK, "error.html", "0")
	}
}

//EchoGameWriteView is shows game write view and insert the written content into the database
func EchoGameWriteView(c echo.Context) error {
	if c.Request().Method == "POST" {
		resgame := c.FormValue("game")
		restype := c.FormValue("type")
		resroot := c.FormValue("root")
		rescontent := c.FormValue("ir1")
		var insertstring = "INSERT INTO game_view (Game, Type, Root, Content) VALUES (" + "'" + resgame + "'" + "," + "'" + restype + "'" + "," + "'" + resroot + "'" + "," + "'" + rescontent + "'" + ");"
		InsertQuery(db1, insertstring)
		http.Redirect(c.Response(), c.Request(), "/menu/?Handler=g_main", http.StatusFound)
	} else {
		return c.Render(http.StatusOK, "game_write.html", "0")
	}
	return c.String(0, "ERROR")
}
