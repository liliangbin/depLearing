package model

import (
	"testing"
	"os"
	"fmt"
	"golang_tes/depLearing/core"
	"log"
	"encoding/json"
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

	var schoolUnit SchoolUnit
	db.Where("id=3").First(&schoolUnit)
	fmt.Println(schoolUnit.UnitName)

	var studentInfos StudentInfo
	db.Where("id = 3").First(&studentInfos)
	fmt.Println(studentInfos.StudentName)

}

func TestWeekClass_InitWeekClass3(t *testing.T) {

	sdf := "{'yb_userid':'8574001','yb_username':'李良彬','yb_usernick':'李良彬','yb_sex':'M','yb_money':'784','yb_exp':'541','yb_userhead':'http://img02.fs.yiban.cn/8574001/avatar/user/200','yb_schoolid':'34039','yb_schoolname':'中国石油大学（华东）','yb_realname':'李良彬','yb_studentid':'1607040211','yb_identity':'学生'}"

	YbInfo := new(YBUserInfo)
	if err := json.Unmarshal([]byte(sdf), &YbInfo); err != nil {
		fmt.Println("解析出错")
		panic(err.Error())
	}

	fmt.Println(YbInfo.YbStudentid)
}
