package main

import (
	"github.com/codegangsta/negroni"
	"net/http"
	"fmt"
	"golang_tes/depLearing/service"
	"os"
	"golang_tes/depLearing/core"
	"log"
	"golang_tes/depLearing/model"
	"golang_tes/depLearing/handlers"
)

func test(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "test for zhe seond test")
}

func init() {

	err := core.ParseConf("config.json")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("请配置config.json")
			os.Exit(0)
		}
		panic(err)

	}

}

func main() {
	initDB()
	router := NewRouter()
	router.HandleFunc("/test", test)
	router.HandleFunc("/login", service.LoginHandler)
	router.HandleFunc("/index", handlers.GetClassInfo)
	router.Handle("/resource", negroni.New(negroni.HandlerFunc(service.ValidateTokenMiddleWare), negroni.WrapFunc(test)))

	n := negroni.Classic()
	n.Use(negroni.NewLogger())
	n.UseHandler(router)
	n.Run(":3000")
}

func initDB() {
	db, err := core.GetSqlConn()
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	var feedback bool

	db.AutoMigrate(model.ClassInfo{})
	db.AutoMigrate(model.StudentInfo{})
	db.AutoMigrate(model.SchoolCourse{})

	db.AutoMigrate(model.SchoolUnit{})

	db.AutoMigrate(model.SchoolTeacher{})

	fmt.Println(feedback)

}
