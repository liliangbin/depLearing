package core

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

func GetSqlConn() (*gorm.DB, error) {
	db :=Conf.Db
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=UTC",
		db.Username, db.Password, db.Address, db.Port, db.Dbname)
	return gorm.Open("mysql", args)
}

func GetMysqlConn() (*sql.DB ,error)  {


	db :=Conf.Db

	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=UTC",
		db.Username, db.Password, db.Address, db.Port, db.Dbname)
	return sql.Open("mysql", args)
}