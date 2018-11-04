package service

import (
	"golang_tes/depLearing/model"
	"fmt"
	"strconv"
	"golang_tes/depLearing/core"
	"log"
	"strings"
)

/*直接计算当前周次，查询当前或是以后的周次。大于18周的时候我们就不再提供查询*/

/*默认周次使用*/

/*service 的设计原则*/

func GetClassInfo(student_id string, class_grade string, week int) interface{} {

	db, err := core.GetSqlConn()
	if err != nil {

		log.Fatal(err.Error())

	}

	defer db.Close()

	var classes []model.ClassInfo

	fmt.Println(student_id , class_grade)
	db.Where("student_id=? and class_grade=?", student_id, class_grade).Find(&classes)


	week_class := new(model.WeekClass)

	for _, class := range classes {

		fmt.Println(class.ClassName, class.ClassHours, len(class.ClassHours), class.ClassWeeks, len(class.ClassWeeks), class.ClassTimes)
		s := class.ClassWeeks
		strs := strings.Split(s, ",")

		/*这个地方我们先把他按照“，”分开，后来我们还需要重新的再按照短横线分开一次*/
		for _, str := range strs {

			//15  两位数的周次剩下的就为单周的数据

			if len(str) > 2 {

				st := strings.Split(str, "-")

				if len(st) != 2 {
					fmt.Println(st, "str====>")
					panic(err.Error())
				}
				low, err1 := strconv.Atoi(st[0])
				high, err := strconv.Atoi(st[1])

				if err != nil || err1 != nil {
					panic(err1.Error())
				}

				if week <= high && week >= low {

					initDayCourses(class, week_class)

				}
			} else {
				//单周计算，假如说达到单周计算的周期那么就加上去。
				tmp, err := strconv.Atoi(str)

				if err != nil {
					panic(err.Error())
				}

				if tmp == week {

					initDayCourses(class, week_class)
				}

			}

		}

	}
	fmt.Println(week_class)
	return week_class

}

func initDayCourses(user model.ClassInfo, week_class *model.WeekClass) {

	/*这个地方开始存储位置*/
	day_class := user.ClassTimes    //星期几的课表信息
	week_day := day_class[:1]       //星期几
	week_day_class := day_class[1:] //直接切割
	fmt.Println(week_day + "\n" + week_day_class)
	this_week, err := strconv.Atoi(week_day)
	if err != nil {
		panic(err.Error())
	}
	for i := 0; i < len(week_day_class); i = i + 2 {

		tmp := week_day_class[i : i+2]

		course_time, err := strconv.Atoi(tmp)
		if err != nil {
			panic(err.Error())
		}
		//获取这个老师的信息

		teacher := GetTeacherInfoByteacherNum(user.ClassTeacher)
		//从零开始。以前是从
		week_class.InsertClass(this_week, course_time-1, user.ClassName, user.ClassPosition, teacher.TeacherName)

	}
}

