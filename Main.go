package main

import (
	"github.com/codegangsta/negroni"
	"net/http"
	"fmt"
	"os"
	"golang_tes/depLearing/core"
	"log"
	"golang_tes/depLearing/model"
	"golang_tes/depLearing/handlers"
	"golang_tes/depLearing/utils"
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
	router.HandleFunc("/oauth", utils.OAuthHandler)
	router.HandleFunc("/class", handlers.GetCurrentClassInfo)
	router.HandleFunc("/class_week", handlers.GetClassInfoByWeekDay)
	router.HandleFunc("/class_grade", handlers.GetClassInfoByWeekAndGrade)
	router.Handle("/resource", negroni.New(negroni.HandlerFunc(utils.ValidateTokenMiddleWare), negroni.WrapFunc(test)))

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
