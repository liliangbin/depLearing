package main

import (
	"github.com/codegangsta/negroni"
	"fmt"
	"os"
	"golang_tes/depLearing/core"
	"log"
	"golang_tes/depLearing/model"
	"golang_tes/depLearing/handlers"
	"golang_tes/depLearing/utils"
)



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
	router.HandleFunc("/oauth", utils.OAuthHandler)
	router.Handle("/class",negroni.New(negroni.HandlerFunc(utils.ValidateTokenMiddleWare),negroni.WrapFunc(handlers.GetCurrentClassInfo)))
	router.Handle("/class_week",negroni.New(negroni.HandlerFunc(utils.ValidateTokenMiddleWare),negroni.WrapFunc(handlers.GetClassInfoByWeekDay)))
	router.Handle("/class_grade",negroni.New(negroni.HandlerFunc(utils.ValidateTokenMiddleWare),negroni.WrapFunc(handlers.GetClassInfoByWeekAndGrade)))

	n := negroni.Classic()
	n.Use(negroni.NewLogger())
	n.UseHandler(router)
	n.Run(":8083")
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
