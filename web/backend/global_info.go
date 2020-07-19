package backend

var db1 = dbInfo{"root", "1463", "localhost:3306", "mysql", "golang_web"}

type dbInfo struct {
	user     string
	pwd      string
	url      string
	engine   string
	database string
}

type notice_board_view struct {
	No 		int //int
	Title 	string
	Writer 	string
	Content string
	Date 	string //date
	Click 	int //int
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
	}
	return reshandle
}