package handlers

import (
	"net/http"
	"fmt"
)

func Hand(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "we can handled it ")
}

func Index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "hello world")
}

func Name(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "hello name")
}

func Wu()  {
	fmt.Println("hahhhh")
}