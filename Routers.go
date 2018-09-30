package main

/*
这个文件里面的routers是主要用于最普通的接口。需要授权的接口我们放在另外的地方来。以后可以从头来操作*/
import (
	"github.com/gorilla/mux"
	"golang_tes/depLearing/model"
	"golang_tes/depLearing/handlers"
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