package model

import (

	_ "github.com/go-sql-driver/mysql"
	"github.com/hb-go/gorm"

	. "sekiro_echo/conf"
	"github.com/hb-go/echo-web/model/orm"
	"github.com/hb-go/echo-web/module/log"
)

var db *gorm.DB

func DB() *gorm.DB {
	if db == nil {
		log.Debugf("Model NewDB")

		newDb, err := newDB()
		if err != nil {
			panic(err)
		}
		newDb.DB().SetMaxIdleConns(10)
		newDb.DB().SetMaxOpenConns(100)

		newDb.SetLogger(orm.Logger{})
		newDb.LogMode(true)
		db = newDb
	}

	return db
}

func newDB() (*gorm.DB, error) {
	sqlConnection := Conf.DB.UserName + ":" + Conf.DB.Pwd + "@tcp(" + Conf.DB.Host + ":" + Conf.DB.Port + ")/" + Conf.DB.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", sqlConnection)
	if err != nil {
		return nil, err
	}
	return db, nil
}
