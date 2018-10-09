package model

type ClassInfo struct {
	ID int 	 `gorm:"primary_key 	AUTO_INCREMENT"`
	ClassGrade string
	ClassName string
	ClassUnit string
	ClassCredits string
	ClassHours string
	ClassTeacher string
	ClassTimes string
	ClassWeeks string
	ClassPosition string
	ClassProperties string
	ClassCategory string
	ClassTeachFor string

}
