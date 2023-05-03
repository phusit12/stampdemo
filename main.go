package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/phusit12/stampdemo.git/routes"
)

// การรีพาย stuc

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello World !")
	})

	uGroup := e.Group("/users")
	routes.NewUserRoute(uGroup)
	e.Logger.Fatal(e.Start(":80")) //ตัว log จะ error จาก e.Start
}
