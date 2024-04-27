package main

import (
	"ginDemo/src/servicce"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", servicce.SayHello)
	r.POST("/add", servicce.AddBodyParam)
	r.PUT("/user", servicce.AddUser)
	r.GET("/user", servicce.SelectByAge)
	r.Run() // listen and serve on 0.0.0.0:8080
}
