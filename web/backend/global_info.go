package backend

var db1 = dbInfo{"root", "1463", "localhost:3306", "mysql", "golang_web"}
var id string = "rodvkf72"
var pw string = "1463"
var header string = "frontend/header.html"
var footer string = "frontend/footer.html"
var leftside string = "frontend/leftside.html"

type dbInfo struct {
	user     string
	pwd      string
	url      string
	engine   string
	database string
}

type notice_board_view struct {
	No      int //int
	Title   string
	Writer  string
	Content string
	Date    string //date
	Click   int    //int
	Maxno	int
}

type notice_board_total_no struct {
	Total_no int
}

type project_view struct {
	No      string
	Title   string
	Content string
	Root    string
}

type game_view struct {
	No   int
	Name string
	Type string
	Root string
}

type hostname struct {
	Name string
}

func static(handle string) string {
	var reshandle string
	switch handle {
	case "n_main":
		reshandle = "1"
	case "n_content":
		reshandle = "2"
	case "n_write":
		reshandle = "3"

	case "p_main":
		reshandle = "11"
	case "p_content":
		reshandle = "12"
	case "p_write":
		reshandle = "13"

	case "g_main":
		reshandle = "21"
	case "g_content":
		reshandle = "22"
	case "g_write":
		reshandle = "23"
	}
	return reshandle
}
