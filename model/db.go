package model

import (

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	. "sekiro-echo/conf"
)

var db *gorm.DB

//DB get db connection
func DB() *gorm.DB {
	if db == nil {

		newDb, err := newDB()
		if err != nil {
			panic(err)
		}
		newDb.DB().SetMaxIdleConns(10)
		newDb.DB().SetMaxOpenConns(100)

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
