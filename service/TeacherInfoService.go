package service

import (
	"golang_tes/depLearing/core"
	"fmt"
	"golang_tes/depLearing/model"
)

func GetTeacherInfoByteacherNum(teacherNum string) *model.SchoolTeacher {

	teacher :=new(model.SchoolTeacher)
	db, err := core.GetMysqlConn()

	defer db.Close()
	if err != nil {
		panic(err.Error())
	}

	err = db.QueryRow("select * from school_teacher where teacher_id = ?", teacherNum).Scan(&teacher.ID,&teacher.TeacherName,&teacher.TeacherId)

	if err != nil {

		panic(err.Error())
	}
	fmt.Println(teacher.TeacherName)

	return teacher
}
