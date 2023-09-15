package main

import (
	"calculator/rpn"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.POST("/calculate", calculate)
	e.Logger.Fatal(e.Start(":1323"))
}

func calculate(c echo.Context) error {
	param := c.QueryParam("equation")
	return c.String(http.StatusOK, strings.Join(rpn.ConvertStrToRpn(param), " "))
}
