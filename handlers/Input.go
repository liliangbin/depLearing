package handlers

import (
	"net/http"
	"fmt"
	"golang_tes/depLearing/model"
)

func Hand(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "we can handled it ")
}

func Index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "hello world")
}

func Name(w http.ResponseWriter, r *http.Request) {
	msg := model.FeedbackMessages{
		UserName: "liliangbin",
		Message:  "ha er 啊",
	}
	msg.Create()
	fmt.Fprint(w, "hello name")
}

func Wu()  {
	fmt.Println("hahhhh")
}