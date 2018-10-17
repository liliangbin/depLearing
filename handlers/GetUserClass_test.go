package handlers

import (
	"testing"
	"os"
	"fmt"
	"golang_tes/depLearing/core"
	"time"
	"strconv"
)

func init() {

	err := core.ParseConf("../config.json")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("请配置config.json")
			os.Exit(0)
		}
		panic(err)
	}
}

func TestGet_User_Class(t *testing.T) {

	GetClass()

	tim := time.Now()
	fmt.Println(tim.ISOWeek())
	timw := time.Date(2018, 9, 8, 1, 1, 1, 1, time.UTC)
	fmt.Println(timw.ISOWeek())

	fmt.Println(strconv.Atoi("01"))
}
