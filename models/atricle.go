package models

import (
	"gorm.io/gorm"
	"log"
)

type Article struct {
	Model

	TagID      int    `json:"tag_id" gorm:"index"`
	Tag        Tag    `json:"tag"`
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func ExistsById(id int) bool {
	var article Article
	log.Println(id)
	db.Select("id").Where("id = ?", id).First(&article)
	return article.ID > 0
}
func SelectById(id int) (article Article) {
	var tag Tag
	db.Where("id = ?", id).First(&article)
	//db.Model(&article).Related(&article.Tag)
	// 关联查询
	db.Model(&article).Association("Tag").Find(&tag)
	article.Tag = tag
	//log.Println(tag)
	return
}

func GetArticleTotal(maps interface{}) (count int64) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	// preLoad等价于提前把Tag查出来，之后可以直接注入到article内部
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}
func EditArticle(id int, data interface{}) bool {
	// 这里用Model先定义了是什么表的感觉
	db.Model(&Article{}).Where("id = ?", id).Updates(data)
	return true
}
func AddArticle(data map[string]interface{}) bool {
	//v表示一个接口值，I表示接口类型。这个实际就是Golang中的类型断言
	//用于判断一个接口值的实际类型是否为某个类型，或一个非接口值的类型是否实现了某个接口类型
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
	return "blog_article"
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
