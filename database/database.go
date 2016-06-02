package db

import (
	"fmt"
	"log"

	r "github.com/dancannon/gorethink"
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

//GetSession gets the current DB session
func GetSession() *r.Session {
	return session
}
