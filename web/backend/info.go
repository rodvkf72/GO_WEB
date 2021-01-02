package backend

//DB 정보나 구조체를 정의
var db1 = dbInfo{"root", "1463", "localhost:3306" /*docker : "172.30.1.19:3306"*/, "mysql", "golang_web"}
var id string = "rodvkf72"
var pw string = "1463"
var head string = "frontend/head.html"
var header string = "frontend/header.html"
var footer string = "frontend/footer.html"
var leftside string = "frontend/leftside.html"

/*
데이터베이스 접속에 필요한 정보를 구조체로 정의
*/
type dbInfo struct {
	user     string
	pwd      string
	url      string
	engine   string
	database string
}

/*
Select문을 Interface로 구현
*/
/*
type Interf interface {
	SelectQuery(db dbInfo, query string) []notice_board_view
}
*/

/*
게시판 목록 출력에 필요한 데이터를 구조체로 정의 (일부 변수는 없어도 되긴 함)
*/
type notice_board_view struct {
	No      int //int
	Title   string
	Writer  string
	Content string
	Date    string //date
	Click   int    //int
	Maxno   int
	Cookie  string
}

/*
게시판 글 클릭 시 출력에 필요한 데이터를 구조체로 정의
*/
type notice_board_content_view struct {
	No         int
	Title      string
	Writer     string
	Content    string
	Date       string
	Click      int
	Maxno      int
	Minno      int
	Minusno    int
	Minustitle string
	Plusno     int
	Plustitle  string
}

/*
프로젝트 목록 출력에 필요한 데이터를 구조체로 정의
*/
type project_view struct {
	No      string
	Title   string
	Content string
	Root    string
}

/*
게임 목록 출력에 필요한 데이터를 구조체로 정의
*/
type game_view struct {
	No      int
	Game    string
	Type    string
	Root    string
	Content string
	Cookie  string
}

type baekjoon_view struct {
	No     int
	Title  string
	Cookie string
}

type baekjoon_content_view struct {
	No      int
	Title   string
	Content string
}

/*
테스트용 구조체
*/
type hostname struct {
	Name string
}

/*
컨트롤러 패턴을 적용하기 위한 핸들러 함수
*/
func static(handle string) string {
	var reshandle string

	switch handle {
	case "n_main": //게시판
		reshandle = "1"
	case "n_content":
		reshandle = "2"
	case "n_write":
		reshandle = "3"

	case "p_main": //프로젝트
		reshandle = "11"
	case "p_content":
		reshandle = "12"
	case "p_write":
		reshandle = "13"

	case "g_main": //게임
		reshandle = "21"
	case "g_content":
		reshandle = "22"
	case "g_write":
		reshandle = "23"

	case "b_main": //백준
		reshandle = "31"
	case "b_content":
		reshandle = "32"
	case "b_write":
		reshandle = "33"
	case "b_search":
		reshandle = "34"
	}

	return reshandle
}
