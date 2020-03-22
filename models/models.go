package models

import (
	"fmt"
	"gin-example/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int ` json:"created_on"`
	ModifiedOn int ` json:"modified_on"`
}

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)
	sec, err := setting.Cfg.GetSection("database")

	if err != nil {
		log.Fatal("fail to get section database")
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()

	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defauleTableName string) string {
		return tablePrefix + defauleTableName
	}

	db.SingularTable(true)
	db.LogMode(true)

	db.DB().SetConnMaxLifetime(10)
	db.DB().SetMaxOpenConns(100)

}

func CloseDB() {
	defer db.Close()

}
