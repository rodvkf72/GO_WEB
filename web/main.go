package main

import (
	"fmt"
	"log"
	"net/http"
	"web/backend"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fs := http.FileServer(http.Dir("./frontend/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", backend.Host_Index)
	http.HandleFunc("/check/", backend.Login_Check)
	http.HandleFunc("/main/", backend.User_Index)
	http.HandleFunc("/fail/", backend.Fail)

	http.HandleFunc("/notice_board/", backend.Request_Handler)
	http.HandleFunc("/notice_board_contents/", backend.Request_Handler)
	http.HandleFunc("/notice_board_write/", backend.Request_Handler)

	http.HandleFunc("/project/", backend.Request_Handler)
	http.HandleFunc("/project_contents/", backend.Request_Handler)
	

	log.Println("Listening on : 9090...")
	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		log.Fatal("ListenAndServer : ", err)
	} else {
		fmt.Println("ListenAndServer Started! -> Port(9000)")
	}
}
