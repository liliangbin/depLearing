package model

import (
	"testing"
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
	"net/url"
	"encoding/json"
)

func TestInterest(t *testing.T) {
	class := make([]Class, 12)
	fmt.Println("sfdfasd")
	class[0] = Class{"4w", "5345", "5345"}

	fmt.Println(class[0].ClassTeacher)

	fmt.Println("fsfsd")

}

func TestWeekClass_InsertClass(t *testing.T) {
	weekDay := new(WeekClass)
	weekDay.InsertClass(3, 4, "sdfsd", "sdfasd", "sdfasd")

	fmt.Println(weekDay)
}

func TestWeekClass_InitWeekClass2(t *testing.T) {

	//appname := "石大请销假"

	//	vq := "c21f263641f5b94c0b20f9ec5d2ad13fa8aee3bd3b2e9a3d911daaa4b393d1e06bd75fde3c1701f590c1fa67afb978677d2c58b2acf0ab55a75d739e0a60a5f882758570d803376281c03858ceda4ca9d9259262baf552e4aa57077ff27c0694a656f2870d41ea9d844140c9a98a83a0754cff37c94f9a97d7aa45c92bce23af85a0ef5f92cceea4bfc793d0f54fd74af6f71de63157891d5814fa9b3ff43398181b20e66c936b4111eb57b30fbc311648e2e9652c4659703638d3661d4580afa3934bdd47f8fa1df58b8e8315336995b3e0568195ff31e777011046e5ebdf672303605c8139470ad2e8107c31ec8dce"
	//resp, err := http.Get("http://localhost:8080/common/user/getStuIdByVq?appName=" + appname + "&vq=" + vq)

	url := "http://localhost:8080/common/user/getStuIdByVq"
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, strings.NewReader("appName=石大请销假"))
	if err != nil {

		panic(err.Error())
	}

	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Charset", "GBK,utf-8;q=0.7,*;q=0.3")
	req.Header.Set("Accept-Encoding", "gzip,deflate,sdch")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")
	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

}

func TestWeekClass_InsertClass2(t *testing.T) {

	vq := "c21f263641f5b94c0b20f9ec5d2ad13fa8aee3bd3b2e9a3d911daaa4b393d1e06bd75fde3c1701f590c1fa67afb978677d2c58b2acf0ab55a75d739e0a60a5f882758570d803376281c03858ceda4ca9d9259262baf552e4aa57077ff27c0694a656f2870d41ea9d844140c9a98a83a0754cff37c94f9a97d7aa45c92bce23af85a0ef5f92cceea4bfc793d0f54fd74af6f71de63157891d5814fa9b3ff43398181b20e66c936b4111eb57b30fbc311648e2e9652c4659703638d3661d4580afa3934bdd47f8fa1df58b8e8315336995b3e0568195ff31e777011046e5ebdf672303605c8139470ad2e8107c31ec8dce"

	resp, err := http.PostForm("http://localhost:8080/common/user/getStuIdByVq",
		url.Values{"key": {"Value"}, "id": {"123"}, "appName": {"石大请销假"}, "vq": {vq}})

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

	YBUserINfo := string(body)
	fmt.Println(YBUserINfo)

	df := new(YBUserInfo)

	if err := json.Unmarshal(body, &df); err != nil {
		fmt.Println("fsdf")
		panic(err.Error())
	}

	fmt.Println(df.YbStudentid)
}

func TestWeekClass_InitWeekClass4(t *testing.T) {

	dfdf := Rii{"2323", "werwerw", "sdfsdf"}

	indff, err := json.Marshal(dfdf)
	if err != nil {

		fmt.Println("加入出错")
		panic(err.Error())
	}
	fmt.Println(string(indff))
	str :="sfdfasdfasff"
	fmt.Println(str)
	df := new(Rii)

	if err := json.Unmarshal(indff, &df); err != nil {
		fmt.Println("fsdf")
		panic(err.Error())
	}

	fmt.Println(df.Name)
}
