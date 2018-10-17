package model

import (
	"golang_tes/depLearing/core"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

)

type FeedbackMessages struct {
	gorm.Model
	UserName string
	Message  string
	ImgUrl   string
}

func (feedback *FeedbackMessages) Create() {

	db, err := core.GetSqlConn()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Create(feedback)
}

func (feedback *FeedbackMessages) insert(username string) {

	feedback.UserName = username

}
