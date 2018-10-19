package model

type ClassInfo struct {
	ID                int `gorm:"primary_key 	AUTO_INCREMENT"`
	ClassGrade        string //开课学年
	ClassName         string //开课名字
	ClassUnit         string //开课所属单位
	ClassCredits      string //学生分
	ClassHours        string //学生学时
	ClassTeacher      string //上课老师
	ClassTimes        string //上课时间
	ClassWeeks        string //上课周次
	ClassPosition     string //上课位置
	ClassProperties   string //课程属性
	ClassCategory     string //素质课课程类别
	ClassTeachFor     string //授课班级
	StudentID         string //学号
	ClassCourseNumber string //课程号
}

type SchoolUnit struct {
	ID       int `gorm:"primary_key 	AUTO_INCREMENT"`
	UnitName string
	UnitId   string
}
type SchoolTeacher struct {
	ID          int `gorm:"primary_key 	AUTO_INCREMENT"`
	TeacherName string
	TeacherId   string
}

type SchoolCourse struct {
	ID           int `gorm:"primary_key 	AUTO_INCREMENT"`
	CourseName   string
	CourseNumber string
}

type StudentInfo struct {
	ID                  int `gorm:"primary_key 	AUTO_INCREMENT"`
	StudentName         string
	StudentId           string
	StudentCollege      string //学生学院
	StudentProfessional string //学生专业
	StudentClass        string //学生班级
}


