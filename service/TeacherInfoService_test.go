package service

import (
	"testing"
	"fmt"
)

func TestGetTeacherInfoByteacherNum(t *testing.T) {

	dfd:=GetTeacherInfoByteacherNum("19900017")
	fmt.Println(dfd)
}
