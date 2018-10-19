package service

import (
	"golang_tes/depLearing/core"
	"fmt"
	"golang_tes/depLearing/model"
)

func GetStuInfoById(stuId string) *model.StudentInfo {

	stu :=new(model.StudentInfo)
	db, err := core.GetMysqlConn()

	if err != nil {
		panic(err.Error())
	}

	err = db.QueryRow("select * from student_infos where student_id = ?", stuId).Scan(&stu.ID, &stu.StudentName, &stu.StudentId, &stu.StudentCollege, &stu.StudentProfessional, &stu.StudentClass)

	if err != nil {

		panic(err.Error())
	}
	fmt.Println(stu.StudentClass)

	return stu
}
