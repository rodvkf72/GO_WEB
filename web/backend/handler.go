package backend

import (
	"github.com/labstack/echo"
)

//EchoRequestHandler is call the static function in the 'info.go' file with the Hnadler value passed in the GET method, and call the corresponding function with the returned value.
func EchoRequestHandler(c echo.Context) error {
	reshandler := c.FormValue("Handler")
	shandle := static(reshandler)

	switch shandle {
	case "1":
		return EchoNoticeboardIndex(c)
	case "2":
		return EchoNoticeboardContentView(c)
	case "3":
		return EchoNoticeboardWriteView(c)

	case "11":
		return EchoProjectIndex(c)
	case "12":
		return EchoProjectContentView(c)
	case "13":
		return EchoProjectWriteView(c)

	case "21":
		return EchoGameIndex(c)
	case "22":
		return EchoGameContentView(c)
	case "23":
		return EchoGameWriteView(c)

	case "31":
		return EchoBaekjoonIndex(c)
	case "32":
		return EchoBaekjoonContentView(c)
	case "33":
		return EchoBaekjoonWriteView(c)
	case "34":
		return EchoBaekjoonSearch(c)
	}
	return c.String(0, "ERROR")
}
