package main

import (
	"net/http"

	"github.com/NagatsukiSep/asset-manager-server/subpkg"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, subpkg.Hello())
	})

	e.Logger.Fatal(e.Start(":1323"))
}
