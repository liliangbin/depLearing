package handlers

import (
	"net/http"
	"golang_tes/depLearing/service"
	"fmt"
	"encoding/json"
	"strconv"
	"golang_tes/depLearing/utils"
)

func GetCurrentClassInfo(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	token := r.Form["token"][0]
	userInfo,err :=utils.GetStuInfoByToken(token)
	utils.CheckError(err)
	student_id :=userInfo.YbStudentid
	fmt.Println("studdent_id",student_id)

	week := utils.GetWeeks()
	fmt.Println(week,"curent weeks ====>")
	initd := service.GetClassInfo(student_id, "2018-2019-1", week)

	service.AllowCORS(w)
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(initd); err != nil {
		panic(err)
	}

}

func GetClassInfoByWeekDay(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	token := r.Form["token"][0]
	week := r.Form["week"][0]
	userInfo,err :=utils.GetStuInfoByToken(token)
	utils.CheckError(err)
	student_id :=userInfo.YbStudentid

	fmt.Println(student_id+"student_id")

	fmt.Println(student_id)
	we, err := strconv.Atoi(week)
	if err != nil {
		panic(err.Error())
	}

	//week := service.GetWeeks()
	fmt.Println(week)
	initd := service.GetClassInfo(student_id, "2018-2019-1", we)

	service.AllowCORS(w)
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(initd); err != nil {
		panic(err)
	}

}

func GetClassInfoByWeekAndGrade(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	token := r.Form["token"][0]
	userInfo,err :=utils.GetStuInfoByToken(token)
	utils.CheckError(err)
	student_id :=userInfo.YbStudentid
	fmt.Println("student_id",student_id)
	week := r.Form["week"][0]
	grade := r.Form["grade"][0]

	fmt.Println(student_id)
	we, err := strconv.Atoi(week)
	if err != nil {
		panic(err.Error())
	}

	//week := service.GetWeeks()
	fmt.Println(week)
	initd := service.GetClassInfo(student_id, grade, we)

	service.AllowCORS(w)
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(initd); err != nil {
		panic(err)
	}

}

