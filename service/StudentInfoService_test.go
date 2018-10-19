package service

import (
	"testing"
	"fmt"
)

func TestGetInfoByStuId(t *testing.T) {
	ff:=GetStuInfoById("1607040211")
	fmt.Println(ff)
}
