package backend

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os/exec"
	"strings"

	"github.com/labstack/echo"
)

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'
	}
}

// Copy the values from channel 'src' to channel 'dst',
// removing those divisible by 'prime'.
func filter(src <-chan int, dst chan<- int, prime int) {
	for i := range src { // Loop over values received from 'src'.
		if i%prime != 0 {
			dst <- i // Send 'i' to channel 'dst'.
		}
	}
}

// The prime sieve: Daisy-chain filter processes together.
func sieve() {
	ch := make(chan int) // Create a new channel.
	go generate(ch)      // Start generate() as a subprocess.
	for {
		prime := <-ch
		fmt.Print(prime, "\n")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}

func sendChannel(ch chan<- string, code string) {
	imp := "package main \n\n import \"fmt\" \n\n func main() {\n fmt.Println(\"" + code + "\")\n }"
	ch <- imp
}

func receiveChannel(ch <-chan string, file string) string {
	data := <-ch
	var bytes []byte
	bytes = []byte(data)

	filename := strings.Replace(file, ".", "", -1)
	err := ioutil.WriteFile("./webCompile/"+filename+".go", bytes, 0644)
	check(err)

	exec.Command("D:/GOWRK/src/GO_WEB/web/webCompile/build.bat", filename).CombinedOutput()

	//out, _ := exec.Command("D:/GOWRK/src/GO_WEB/web/webCompile/go run test.go").Output()
	//cmd := exec.CommandContext("D:/GOWRK/src/GO_WEB/web/webCompile", "go run test.go")

	//cmd := exec.Command("D:/GOWRK/src/GO_WEB/web/webCompile/" + filename + ".exe")
	cmd := exec.Command("D:/GOWRK/src/GO_WEB/web/webCompile/" + filename + ".exe")

	byteOutput, _ := cmd.Output()
	data = string(byteOutput)

	return data
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func CodeAjax(c echo.Context) error {
	rescode := c.FormValue("code")

	channel := make(chan string, 1)
	sendChannel(channel, rescode)

	//var file string = string(getOutboundIP())
	var file string

	addrs, _ := net.InterfaceAddrs()

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			file = ipnet.IP.String()
		}
	}

	result := receiveChannel(channel, file)

	fmt.Println(result)
	return c.String(http.StatusOK, result)
}

/*
func getOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
*/
