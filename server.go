package main

import (
	"github.com/labstack/echo/v4"
	"job-backend/controllers"
)

func main() {
	e := echo.New()
	user := e.Group("/user")
	user.POST("/register", controllers.Register)
	e.Logger.Fatal(e.Start(":3500"))
}
