package model

import "fmt"

type Class struct {
	ClassName     string
	ClassPosition string
	ClassTeacher  string
}

type Day struct {
	Course [12]Class
}

type WeekClass struct {
	Days [7]Day
}

type WeekClassInterface interface {
	InitWeekClass()
	InsertClass(weekDay int, courseTime int, courseName string, coursePosition string, courseTeacher string)
}

func (weekClass *WeekClass) InitWeekClass() {

	weekClass.Days[0].Course[7].ClassTeacher = "hello"

}

func (weekClass *WeekClass) InsertClass(weekDay int, courseTime int, courseName string, coursePosition string, courseTeacher string) {

	weekClass.Days[weekDay].Course[courseTime].ClassName = fmt.Sprintf("%s %s", weekClass.Days[weekDay].Course[courseTime].ClassName, courseName)

	weekClass.Days[weekDay].Course[courseTime].ClassPosition = fmt.Sprintf("%s %s", weekClass.Days[weekDay].Course[courseTime].ClassPosition, coursePosition)
	weekClass.Days[weekDay].Course[courseTime].ClassTeacher = fmt.Sprintf("%s %s", weekClass.Days[weekDay].Course[courseTime].ClassTeacher, courseTeacher)


/*	weekClass.Days[weekDay].Course[courseTime].ClassName = courseName
	weekClass.Days[weekDay].Course[courseTime].ClassPosition = coursePosition
	weekClass.Days[weekDay].Course[courseTime].ClassTeacher = courseTeacher
*/
}
