package service

import (
	"time"
	"fmt"
)

/*返回的是开学的时候距离现在的周次。*/
func GetWeeks() int {

	timStart := time.Date(2018, 9, 8, 1, 1, 1, 1, time.UTC)
	time := time.Now()

	_, weeks1 := timStart.ISOWeek()

	_, weeks2 := time.ISOWeek()
	fmt.Println(weeks1 ,"======>")
	return weeks2 - weeks1

}

/*使用的是week*/
