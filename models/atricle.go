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

func ExistsById(id int) (bool, error) {
	var article Article
	log.Println(id)
	tx := db.Select("id").Where("id = ?", id).First(&article)

	if tx.Error != nil && tx.Error != gorm.ErrRecordNotFound {
		return false, tx.Error
	}

	return article.ID > 0, nil
}
func SelectById(id int) (article Article, err error) {
	var tag Tag
	tx := db.Where("id = ?", id).First(&article)

	if tx.Error != nil {
		return article, tx.Error
	}
	//db.Model(&article_service).Related(&article_service.Tag)
	// 关联查询
	err = db.Model(&article).Association("Tag").Find(&tag)
	if err != nil {
		return article, tx.Error
	}
	article.Tag = tag
	//log.Println(tag)
	return
}

func GetArticleTotal(maps interface{}) (count int64, err error) {
	tx := db.Model(&Article{}).Where(maps).Count(&count)

	return count, tx.Error
}

func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []*Article, err error) {
	// preLoad等价于提前把Tag查出来，之后可以直接注入到article内部
	tx := db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	if tx.Error != nil {

		return articles, tx.Error
	}
	return
}
func EditArticle(id int, data interface{}) error {
	// 这里用Model先定义了是什么表的感觉
	tx := db.Model(&Article{}).Where("id = ?", id).Updates(data)

	if tx.Error != nil {
		return tx.Error
	}
	return nil
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
func DeleteArticle(id int) error {
	tx := db.Where("id = ?", id).Delete(Article{})
	if tx.Error != nil && tx.Error != gorm.ErrRecordNotFound {
		return tx.Error
	}
	return nil
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
