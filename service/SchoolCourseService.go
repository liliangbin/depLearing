package service

import (
	"golang_tes/depLearing/model"
	"golang_tes/depLearing/core"
)

func GetCourseInfoByCourseNum(courseNum string) *model.SchoolCourse  {

	db,err :=core.GetMysqlConn()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	return nil
}
