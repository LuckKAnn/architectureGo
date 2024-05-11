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

func init() {
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatalf("Fatal to load database conf: %v", err)
	}
	dbType := sec.Key("TYPE").String()
	dbName := sec.Key("NAME").String()
	user := sec.Key("USER").String()
	password := sec.Key("PASSWORD").String()
	host := sec.Key("HOST").String()
	tablePrefix := sec.Key("TABLE_PREFIX").String()
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
