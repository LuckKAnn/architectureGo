package models

import (
	"gorm.io/gorm"
	"log"
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

func ExistsById(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)
	return article.ID > 0
}
func SelectById(id int) (article Article) {
	db.Select("id").Where("id = ?", id).First(&article)
	//db.Model(&article).Related(&article.Tag)
	return
}

func GetArticleTotal(maps interface{}) (count int64) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}
func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}
func EditArticle(id int, data interface{}) bool {
	db.Where("id = ?", id).Updates(data)
	return true
}
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
func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(Article{})

	return true
}
func (article Article) TableName() string {
	return "blog_atricle"
}

func (tag *Article) BeforeCreate(db *gorm.DB) error {
	log.Println("Invoke BeforeCreate Method")
	//tx := db.Model(&tag).Where("name = ?", tag.Name).Update("CreatedOn", time.Now().Unix())
	//
	////tx := db.UpdateColumn("CreatedOn", time.Now().Unix())
	//if tx.Error != nil {
	//	log.Fatalf("Invoke BeforeCreate Method Fail, %v", tx.Error)
	//}
	return nil
	//return tx.Error
}

func (tag *Article) BeforeUpdate(db *gorm.DB) error {
	log.Println("Invoke BeforeUpdate Method")
	//tx := db.Model(&tag).Where("name = ?", tag.Name).Update("ModifiedOn", time.Now().Unix())
	//
	////tx := db.UpdateColumn("CreatedOn", time.Now().Unix())
	//if tx.Error != nil {
	//	log.Fatalf("Invoke BeforeCreate Method Fail, %v", tx.Error)
	//}
	return nil
}
