package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/codegangsta/negroni"
	"net/http"
	"fmt"
	"golang_tes/depLearing/service"
)



func test(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "test for zhe seond test")
}

func main() {
	router := NewRouter()
	router.HandleFunc("/test", test)
	router.HandleFunc("/login", service.LoginHandler)

	router.Handle("/resource", negroni.New(negroni.HandlerFunc(service.ValidateToeknMiddleWare), negroni.WrapFunc(test)))

	n := negroni.Classic()
	n.Use(negroni.NewLogger())
	n.UseHandler(router)
	n.Run(":3000")
}
