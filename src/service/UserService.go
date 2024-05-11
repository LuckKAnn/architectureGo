package service

import (
	"context"
	"encoding/json"
	"fmt"
	"ginDemo/src/models"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

type UserController struct {
}

const (
	ETCD_PREFIX = "TEST_ETCD_01_"
)

func AddUser(c *gin.Context) {
	paramName := c.Query("name")
	user := models.User{
		Username: paramName,
		Age:      25,
		Email:    string(rune(rand.Int())) + "@qq.com",
	}
	models.DB.Create(&user)
	marshal, err := json.Marshal(&user)
	if err != nil {
		fmt.Println("序列化错误")
		return
	}
	err = models.InsertData(ETCD_PREFIX+user.Username, string(marshal))

	if err != nil {
		fmt.Println("插入etcd错误")
		return
	}
	c.String(http.StatusOK, "增加数据成功!")
	fmt.Println("user==>", user)
}

func SelectByAge(c *gin.Context) {
	userList := []models.User{}
	// 指针
	models.DB.Where("age>20").Find(&userList)
	c.JSON(http.StatusOK, gin.H{"data": userList})
}

func SelectById(c *gin.Context) {
	//models.EtcdClient.
	// get
	paramUserName := c.Query("name")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := models.EtcdClient.Get(ctx, ETCD_PREFIX+paramUserName)
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		//c.JSON(http.StatusOK, gin.H{"data": ev.Value})
		c.String(http.StatusOK, string(ev.Value))

		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}
}
