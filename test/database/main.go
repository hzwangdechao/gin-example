package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	var (
		err error
	)
	db, err = gorm.Open("mysql", "root:123456@tcp(106.14.214.86:3306)/blog?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defauleTableName string) string {
		return "blog_" + defauleTableName
	}

	db.SingularTable(true)
	db.LogMode(true)

	db.DB().SetConnMaxLifetime(10)
	db.DB().SetMaxOpenConns(100)

}

func CloseDB() {
	defer db.Close()

}

type Model struct {
	Id int `gorm:"primary_key;column:id" json:"id"`
}

type Tag struct {
	Model
	Name string `json:"name"`
}

func main() {
	//BelongTo()
	QueryTwo()
}

func BelongTo() {
	// 默认外键名， 类型名+外键名， 即User+Nickname=UserNickname
	type User struct {
		Nickname string `gorm:"primary_key" json:"nickname"`
		Username string `gorm:"column:username" json:"username"`
	}

	type Article struct {
		Model
		TagId        int    `gorm:"column:tag_id" json:"tag_id"`
		Tag          Tag    `json:"tag"`
		Title        string `gorm:"title" json:"title"`
		Desc         string `gorm:"desc" json:"desc"`
		UserNickname string `gorm:"column:nickname;index;" json:"user_id"`
		User         User `gorm:"foreignkey:UserNickname"`
	}
	var article Article
	db.Preload("User").Where("id=222").First(&article)
	log.Println(article)
}

func QueryTwo() {
	// 从表的外键来源于主表的非主键
	type Tt struct {
		Title    string `gorm:"title"`
		TitleTag string `gorm:"title_tag"`
	}
	type Article struct {
		Model
		TtTitle     string `gorm:"column:tt_title"`
		Desc      string
		Tt Tt `gorm:"association_foreignkey:Title"`
	}
	var article Article
	var title Tt
	db.Where("id=222").First(&article)
	db.Model(&article).Related(&title)
	//db.Preload("TitleInfo").Where("id=222").First(&article)

	fmt.Println(article)
	fmt.Println(title)

}
