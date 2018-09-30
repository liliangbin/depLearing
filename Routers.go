package main

import (
	"github.com/gorilla/mux"
	"./model"
	"./handlers"
)
func NewRouter() *mux.Router {
	router :=mux.NewRouter().StrictSlash(true)
	for _,route := range routes{
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandleFunc)
	}
	return router
}


var routes  = model.Routes{

	model.Router{
		"Index",
		"GET",
		"/",
		handlers.Index,
	},
	model.Router{
		"Name",
		"GET",
		"/name",
		handlers.Name,
	},
	model.Router{
		"Handle",
		"GET",
		"/handle",
		handlers.Hand,
	},
}