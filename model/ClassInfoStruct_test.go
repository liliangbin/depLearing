package model

import (
	"testing"
	"os"
	"fmt"
	"golang_tes/depLearing/core"
	"log"
)

func init() {

	err := core.ParseConf("../config.json")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("请配置config.json")
			os.Exit(0)
		}
		panic(err)
	}
}

func TestWeekClass_InitWeekClass(t *testing.T) {

	db, err := core.GetSqlConn()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	var schoolTeacher SchoolTeacher

	db.Where("teacher_name= ?", "王风华").First(&schoolTeacher)

	fmt.Println(schoolTeacher.TeacherName)

	var schoolCourse SchoolCourse
	db.Where("id=3").First(&schoolCourse)
	fmt.Println(schoolCourse.CourseName)

	log.Fatal("hello world")

	var schoolUnit  SchoolUnit
	db.Where("id=3").First(&schoolUnit)
	fmt.Println(schoolUnit.UnitName)

	var studentInfos StudentInfo
	db.Where("id = 3").First(&studentInfos)
	fmt.Println(studentInfos.StudentName)

}

