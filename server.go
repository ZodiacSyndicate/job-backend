package main

import (
	"fmt"
	"job-backend/controllers"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	user := e.Group("/user")
	user.POST("/register", controllers.Register)
	user.POST("/login", controllers.Login)
	user.GET("/info", controllers.Info)
	user.GET("/list", controllers.List)
	user.GET("/", func(c echo.Context) error {
		fmt.Println(uuid.New().String())
		return nil
	})
	e.Logger.Fatal(e.Start(":3500"))
}
