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
	db.DropTable(&User{}, &Dept{}, &DeptPos{})

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Dept{})
	db.AutoMigrate(&DeptPos{})

}

func Create() {
	var user = &User{
		Nickname: "wdc",
		Username: "wangdechao",
		Sex:      1,
		DeptId:   2902,
	}
	db.Create(&user)

	var dept = &Dept{
		DeptId:   2902,
		DeptName: "yibu",
	}
	db.Create(&dept)

	var depts = &DeptPos{
		DeptId: 2902,
		DeptW:  "hangzhou",
	}
	db.Create(&depts)


}

//type User struct {
//	Nickname        int        `gorm:"TYPE:int(11);NOT NULL;PRIMARY_KEY;INDEX"`
//	Name      string     `gorm:"TYPE: VARCHAR(255); DEFAULT:'';INDEX"`
//	Companies []Company  `gorm:"FOREIGNKEY:UserIds;ASSOCIATION_FOREIGNKEY:Nickname"`
//	CreatedAt time.Time  `gorm:"TYPE:DATETIME"`
//	UpdatedAt time.Time  `gorm:"TYPE:DATETIME"`
//	DeletedAt *time.Time `gorm:"TYPE:DATETIME;DEFAULT:NULL"`
//}

type Company struct {
	Nick     int    `gorm:"primary_key"`
	Industry int    `gorm:"TYPE:INT(11);DEFAULT:0"`
	Name     string `gorm:"TYPE:VARCHAR(255);DEFAULT:'';INDEX"`
	Job      string `gorm:"TYPE:VARCHAR(255);DEFAULT:''"`
	UserIds  int    `gorm:"TYPE:int(11);NOT NULL;INDEX"`
}

type User struct {
	Nickname string `gorm:"cloumn:nickname;primary_key"`
	Username string `gorm:"column:username"`
	Sex      int
	DeptId   int `gorm:"column:dept_id"`
}

type Dept struct {
	DeptId   int    `gorm:"column:dept_id;primary_key"`
	DeptName string `gorm:"column:dept_name"`
	Users    []User `gorm:"foreignkey:DeptId;ASSOCIATION_FOREIGNKEY:DeptId"`
	DeptPos DeptPos `gorm:"foreignkey:DeptId;ASSOCIATION_FOREIGNKEY:DeptName"`
}

type DeptPos struct {
	DeptId int    `gorm:"column:dept_id;primary_key"`
	DeptW  string `gorn:"column:dept_w"`
	DeptName  string `gorn:"column:dept_name"`
}

// foreignkey 指向关联表中的关联字段，  ASSOCIATION_FOREIGNKEY 指向当前表中的主键
func main() {
	Create()
	//var company Company
	//var user User
	var dept Dept

	db.First(&dept, "dept_id=?", "2902")
	//db.Where("dept_id=2902").Preload("Users").Preload("DeptPos").First(&dept)
	//db.Model(&dept).Related(&user)
	//
	db.Model(&dept).Association("Users").Find(&dept.Users)
	fmt.Println(dept)
	//fmt.Println(user)

}
