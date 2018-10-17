package service

import (
	"time"
)

func GetWeeks() int {

	timStart := time.Date(2018, 9, 8, 1, 1, 1, 1, time.UTC)
	time := time.Now()

	_, weeks1 := timStart.ISOWeek()

	_, weeks2 := time.ISOWeek()

	return weeks2 - weeks1
}
