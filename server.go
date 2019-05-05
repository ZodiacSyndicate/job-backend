package main

import (
	"job-backend/controllers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	user := e.Group("/user")
	user.POST("/register", controllers.Register)
	e.Logger.Fatal(e.Start(":3500"))
}
