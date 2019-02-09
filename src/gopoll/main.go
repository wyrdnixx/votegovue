package main

import (
	"fmt"
	//gonfig -> config aus json file lesen
	//"github.com/wyrdnixx/mygoproject/cmd/api/modules"

	// "database/sql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	// _ "github.com/mattn/go-sqlite3"
)

//var AppConfig = modules.Configuration{}

func main() {

	fmt.Printf("Hallo")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/polls", func(c echo.Context) error {
		fmt.Printf("INFO: GET polls ")
		return c.JSON(200, "GET Polls")
	})

	e.PUT("/polls", func(c echo.Context) error {
		fmt.Printf("INFO: PUT polls ")
		return c.JSON(200, "PUT Polls")
	})

	e.PUT("/polls/:id", func(c echo.Context) error {
		fmt.Printf("INFO: PUT polls:id ")
		return c.JSON(200, "UPDATE Poll "+c.Param("id"))
	})

	e.Logger.Fatal(e.Start(":8080"))
}
