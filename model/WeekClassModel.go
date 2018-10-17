package model

type Class struct {
	ClassName     string
	ClassPosition string
	ClassTeacher  string
}

type Day struct {
	Course []Class
}

type WeekClass struct {
	Days []Day
}

func InitWeekClass() *WeekClass {

	weekClass := new(WeekClass)
	weekClass.Days = make([]Day, 7)
	for _, week_day := range weekClass.Days {

		week_day.Course = make([]Class, 12)
	}
	return weekClass
}
