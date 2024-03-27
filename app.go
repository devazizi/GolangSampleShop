package main

import (
	"app/infrastructures"
	"app/router"
	"fmt"

	// "reflect"
	// "fmt"

	echo "github.com/labstack/echo/v4"
)

func main() {
	db_dsn := "gitea:<gitea>@tcp(127.0.0.1:3306)/gitea"
	db := infrastructures.NewMySqlDB(db_dsn)
	query, err := db.Query("SHOW TABLES LIKE 'a%';")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(query)
	fmt.Println(db)
	e := echo.New()
	router.Routes(e)
	// fmt.Println("hello world this is a simple golang application")
	e.Logger.Fatal(e.Start("0.0.0.0:3400"))
}
