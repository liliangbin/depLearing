package model

type StudentInfo struct {
	ID int 	`gorm:"primary_key 	AUTO_INCREMENT"`
	StudentName string
	StudentId string
	StudentCollege string  //学生学院
	StudentProfessional string //学生专业
	StudentClass string //学生班级

}
