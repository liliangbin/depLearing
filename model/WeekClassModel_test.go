package model

import (
	"testing"
	"fmt"
)

func TestInterest(t *testing.T) {
	class := make([]Class ,12)
	fmt.Println("sfdfasd")
	class[0] = Class{"4w","5345","5345"}

	fmt.Println(class[0].ClassTeacher)

	fmt.Println("fsfsd")


}

func TestWeekClass_InsertClass(t *testing.T) {
	weekDay :=new (WeekClass)
	weekDay.InsertClass(3,4,"sdfsd","sdfasd","sdfasd")

	fmt.Println(weekDay)
}

