package main

import (
	"ginDemo/src/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", service.SayHello)
	r.POST("/add", service.AddBodyParam)
	r.PUT("/user", service.AddUser)
	r.GET("/user", service.SelectByAge)
	r.GET("/user/id", service.SelectById)
	r.GET("/user/insert", service.AddUser)

	r.Run() // listen and serve on 0.0.0.0:8080
}
