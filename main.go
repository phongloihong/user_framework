package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/phongloihong/user_framework/router"
)

func main() {
	e := echo.New()
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, //1kb
	}))
	e.Use(middleware.CORS())
	e.Use(middleware.RequestID())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: 5}))

	router.Public(e)
	router.Staff(e)

	fmt.Println("Server listening at 9090")

	e.Start(":9090")
}
