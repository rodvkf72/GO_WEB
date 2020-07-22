package backend

import (
	"net/http"
	"fmt"
)

func Login_Check(w http.ResponseWriter, r *http.Request) {
	//ip 주소 검색
	/*addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String() + "\n")
			}
		}
	}*/
	
	// 아이디 비밀번호 확인
	r.ParseForm()
	if r.Method == "POST" {
		resid := r.FormValue("id")
		respw := r.FormValue("pw")
		if (resid == id) && (respw == pw) {
			fmt.Println("OK")
			http.Redirect(w, r, "/main/", http.StatusFound)
		} else {
			fmt.Println("U R Not Admin !")
			http.Redirect(w, r, "/fail/", http.StatusFound)
		}
	}
}