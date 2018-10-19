package handlers

import (
	"net/http"
	"golang_tes/depLearing/service"
	"fmt"
	"encoding/json"
)

func GetClassInfo(w http.ResponseWriter, r *http.Request) {


	r.ParseForm()

	initd := service.GetClassInfo("1607040211","2018-2019",6)

	service.AllowCORS(w)
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "hello world ")
	if err := json.NewEncoder(w).Encode(initd); err != nil {
		panic(err)
	}

}

/*调用时间太久了，如果有考虑的话可以考虑加入缓存*/