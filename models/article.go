package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	Model
	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreateOn", time.Now().Unix())
	return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil

}

// 通过id判断文章是否存在
func ExistsArticleById(id int) bool {
	var article Article
	db.Select("id").Where("id=?", id).First(&article)
	if article.ID > 0 {
		return true
	}
	return false
}

// 通过map条件查询文章数量
func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

// 通过map条件查询文章详情
func GetArticles(pageNum, pageSize int, maps interface{}) (articles []Article) {
	// 预加载查询
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}

// 通过id查询指定文章
func GetArticle(id int) (article Article) {
	// 这里用到关联查询
	db.Model(&Article{}).Where("id=?", id).First(&article)
	db.Model(&article).Related(&article.Tag)
	return
}

// 通过ID 更新文章
func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id=?", id).Updates(data)
	return true
}

// 添加文章
func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})
	return true
}

// 通过id 删除文章
func DeleteArticle(id int) bool {
	db.Where("id=?", id).Delete(&Article{})
	return true
}
