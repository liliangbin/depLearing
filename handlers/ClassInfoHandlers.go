package handlers

import (
	"golang_tes/depLearing/core"
	"log"
	"golang_tes/depLearing/model"
	"fmt"
	"strings"
	"golang_tes/depLearing/service"
	"strconv"
)

func GetClass() {

	db, err := core.GetSqlConn()
	if err != nil {

		log.Fatal(err.Error())
	}

	defer db.Close()
	studentid := 1607040211
	class_grade := "2018-2019"

	var users []model.ClassInfo

	db.Where("student_id=? and class_grade=?", studentid, class_grade).Find(&users)

	week := service.GetWeeks()

	week_class := model.InitWeekClass()

	for _, user := range users {

		fmt.Println(user.ClassName, user.ClassHours, len(user.ClassHours), user.ClassWeeks, len(user.ClassWeeks))
		s := user.ClassWeeks
		strs := strings.Split(s, ",")

		/*这个地方我们先把他按照“，”分开，后来我们还需要重新的再按照短横线分开一次*/
		for _, str := range strs {

			//15  两位数的周次
			if len(str) > 2 {

				strss := strings.Split(str, "-")

				if len(strss) != 2 {
					fmt.Println(strss)
					panic(err.Error())
				}
				low, err1 := strconv.Atoi(strss[0])
				high, err := strconv.Atoi(strss[1])

				if err != nil || err1 != nil {
					panic(err1.Error())
				}

				if week <= high && week >= low {

					/*这个地方开始存储位置*/

				}
			}
		}

		fmt.Println(week_class.Days[0])
	}

}
