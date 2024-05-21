package models

import (
	"gorm.io/gorm"
	"log"
	"time"
)

// Tag 数据类型
// 方法也要写在这里面吗
type Tag struct {
	Model
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func (tag Tag) TableName() string {
	return "blog_tag"
}

// 这个map是什么呢
func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag, err error) {
	tx := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	if tx.Error != nil && tx.Error != gorm.ErrRecordNotFound {
		return nil, tx.Error
	}
	return
}

func GetTagTotal(maps interface{}) (count int64) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	// 直接返回已经声明过的对象
	return
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ? ", name).First(&tag)

	if tag.ID > 0 {
		return true
	}
	return false
}
func ExistTagById(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	return tag.ID > 0
}
func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})
	return true
}

func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)
	return true
}

func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})
	return true
}
func (tag *Tag) BeforeCreate(db *gorm.DB) error {
	log.Println("Invoke BeforeCreate Method")
	tx := db.Model(&tag).Where("name = ?", tag.Name).Update("CreatedOn", time.Now().Unix())

	//tx := db.UpdateColumn("CreatedOn", time.Now().Unix())
	if tx.Error != nil {
		log.Fatalf("Invoke BeforeCreate Method Fail, %v", tx.Error)
	}
	return tx.Error
}

func (tag *Tag) BeforeUpdate(db *gorm.DB) error {
	log.Println("Invoke BeforeUpdate Method")
	tx := db.Model(&tag).Where("name = ?", tag.Name).Update("ModifiedOn", time.Now().Unix())

	//tx := db.UpdateColumn("CreatedOn", time.Now().Unix())
	if tx.Error != nil {
		log.Fatalf("Invoke BeforeCreate Method Fail, %v", tx.Error)
	}
	return nil
}
