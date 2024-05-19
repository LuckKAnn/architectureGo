package models

import (
	"fmt"
	"ginDemo/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// 加载数据库模型
var db *gorm.DB

// 应该算是baseMode
type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func Setup() {
	var err error
	dbType := setting.DatabaseSetting.Type
	dbName := setting.DatabaseSetting.Name
	user := setting.DatabaseSetting.User
	password := setting.DatabaseSetting.Password
	host := setting.DatabaseSetting.Host
	tablePrefix := setting.DatabaseSetting.TablePrefix
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println(dbType)
	fmt.Println(tablePrefix)

}
