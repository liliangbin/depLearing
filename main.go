package main

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"github.com/codegangsta/negroni"
	"database/sql"
)

func init() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/liliangbin?charset=utf8")
	defer db.Close()
	if err != nil {
		log.Fatal(err, "heihei ......")
	}
}

func main() {
	router := NewRouter()
	n := negroni.Classic()
	n.Use(negroni.NewLogger())
	n.UseHandler(router)
	n.Run(":3000")
}
