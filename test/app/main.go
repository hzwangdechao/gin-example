package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB

func init() {
	var (
		err error
	)
	db, err = gorm.Open("mysql", "root:123456@tcp(106.14.214.86:3306)/test?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Println(err)
	}

	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defauleTableName string) string {
	//	return "blog_" + defauleTableName
	//}

	db.SingularTable(true)
	db.LogMode(true)

	db.DB().SetConnMaxLifetime(10)
	db.DB().SetMaxOpenConns(100)

}

type BaseInfo struct {
	AppId   int    `gorm:"column:appId;primary_key"`
	Cc      string `gorm:"column:cc;primary_key"`
	Version string `gorm:"column:version"`
	Icon    string `gorm:"column:icon"`

	VersionInfo VersionInfo `gorm:"foreignkey:AppId;association_foreignkey:AppId"`
}

type VersionInfo struct {
	AppId   int    `gorm:"column:appId;primary_key"`
	Cc      string `gorm:"column:cc;primary_key"`
	Version string `gorm:"column:version"`

	AppName string `gorm:"column:appName"`
}

func main() {
	Create()

	var base BaseInfo
	var version VersionInfo

	db.Joins("inner join version_info on base_info.appId=version_info.appId and base_info.cc=version_info.cc").Find(&base)

	fmt.Println(base)
	fmt.Println(version)

}

func Create() {
	db.DropTableIfExists(&BaseInfo{}, &VersionInfo{})
	db.AutoMigrate(&BaseInfo{}, &VersionInfo{})

	base1 := BaseInfo{
		AppId:   1,
		Cc:      "cn",
		Version: "1.0",
		Icon:    "http",
	}
	version1 := VersionInfo{
		AppId:   1,
		Cc:      "cn",
		Version: "1.0",
		AppName: "one",
	}
	//version2 := VersionInfo{
	//	AppId:   1,
	//	Cc:      "cn",
	//	Version: "2.0",
	//	AppName: "one",
	//}

	db.Create(&base1)
	db.Create(&version1)
	//db.Create(&version2)

	//db.Debug().Model(&BaseInfo{}).AddForeignKey("`appId` , `cc` , `version`", "version_info(`appId` , `cc` , `version`)", "CASCADE", "CASCADE")

}
