package servicce

import (
	"fmt"
	"ginDemo/src/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
}

func AddUser(c *gin.Context) {
	user := models.User{
		Username: "zhaoliu",
		Age:      25,
		Email:    "111@qq.com",
	}
	models.DB.Create(&user)
	c.String(http.StatusOK, "增加数据成功!")
	fmt.Println("user==>", user)
}

func SelectByAge(c *gin.Context) {
	userList := []models.User{}
	// 指针
	models.DB.Where("age>20").Find(&userList)
	c.JSON(http.StatusOK, gin.H{"data": userList})
}
