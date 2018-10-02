package core

import (
	"testing"
	"log"
	"os"
	"fmt"
)

/*在做功能测试的时候避免出现循环引包*/
/*test用于功能尽量简单的东西*/
func init() {

	err := ParseConf("../config.json")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("请配置config.json")
			os.Exit(0)
		}
		panic(err)
	}
}

func TestGetSqlConn(t *testing.T) {
	log.Println(GetSqlConn())
}
