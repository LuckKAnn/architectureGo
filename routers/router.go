package routers

import (
	"ginDemo/pkg/setting"
	v1 "ginDemo/routers/api/v1"
	"ginDemo/src/service"
	"github.com/gin-gonic/gin"
)

func InitGinServer() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	r.GET("/", service.SayHello)
	r.POST("/add", service.AddBodyParam)
	r.PUT("/user", service.AddUser)
	r.GET("/user", service.SelectByAge)
	r.GET("/user/id", service.SelectById)
	r.GET("/user/insert", service.AddUser)

	apiV1 := r.Group("/api/v1")
	{
		//获取标签列表
		apiV1.GET("/tags", v1.GetTags)
		//新建标签
		apiV1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiV1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiV1.DELETE("/tags/:id", v1.DeleteTag)
	}
	return r
}
