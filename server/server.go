package main

import (
	"fmt"
	"log"
	"net/http"

	r "github.com/dancannon/gorethink"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
)

var (
	session *r.Session
)

func init() {
	var err error

	session, err = r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "test",
		MaxOpen:  40,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}
	if err == nil {
		fmt.Println("Database Connection Successful")
	}
}

func saveStore(c echo.Context) error {
	u := &Supplier{}

	if errA := c.Bind(u); errA != nil {
		return errA
	}

	if errB := r.Table("stores").Insert(u).Exec(session); errB != nil {
		return errB
	}

	return c.JSON(200, u)
}

func main() {

	router := echo.New()

	router.GET("/store", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	router.POST("/store", saveStore)

	router.Run(fasthttp.New(":8000"))
}
