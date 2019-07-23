package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"time"
	"log"
)

func main() {
	e := echo.New()

	ctl := setup()
	e.Use(MiddlewareAPIProfiler())
	e.Use(middleware.Logger())

	e.GET("/metrics", ctl.Metrics)
	e.GET("/sleep/:mean/:std", ctl.Sleep)
	e.GET("/", ctl.Hello)

	e.Logger.Fatal(e.Start(":1323"))
}

func setup() *Controller {
	db, err := sql.Open("mysql", "root:password@tcp(mysql:3306)/prometheus")
	if err != nil {
		panic(err.Error())
	}
	for {
		err := db.Ping();
		if err == nil {
			break
		}
		log.Printf("%v\n", err)
		time.Sleep(time.Second)
	}
	return &Controller{
		Repo: &GreetingRepoMySQL{db: db},
	}
}
