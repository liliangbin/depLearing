package model

type StudentInfo struct {
	StudentName string 	`gorm:"primary_key 	AUTO_INCREMENT"`
	StudentId int
	StudentCollege string
	StudentProfessional string
	StudentClass string
	Classes []ClassInfo `gorm:"many2many:students_classes;"`

}