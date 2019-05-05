package controllers

import (
	"github.com/jinzhu/gorm"
	// mysql driver init
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "hehe:woaicai@tcp(127.0.0.1:3306)/flutter_job?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)
	if !db.HasTable(&User{}) {
		db.CreateTable(&User{})
	}
}
