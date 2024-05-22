package routers

import (
	"ginDemo/pkg/export"
	"ginDemo/pkg/qrcode"
	"ginDemo/pkg/setting"
	"ginDemo/pkg/upload"
	"ginDemo/routers/api"
	v1 "ginDemo/routers/api/v1"
	"ginDemo/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitGinServer() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	r.GET("/auth", api.GetAuth)
	r.POST("/upload", api.UploadImage)
	r.POST("/tags/import", v1.ImportTags)
	// 增加本地静态文件的访问
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

	r.GET("/", service.SayHello)
	r.POST("/add", service.AddBodyParam)
	r.PUT("/user", service.AddUser)
	r.GET("/user", service.SelectByAge)
	r.GET("/user/id", service.SelectById)
	r.GET("/user/insert", service.AddUser)

	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullSavePath()))

	apiV1 := r.Group("/api/v1")

	// 针对某组api开启token校验
	//apiV1.Use(jwt.JWT())
	{
		//获取标签列表
		apiV1.GET("/tags", v1.GetTags)
		r.POST("/tags/export", v1.ExportExcel)
		apiV1.POST("/articles/poster/generate", v1.GenerateArticlePoster)
		//新建标签
		apiV1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiV1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiV1.DELETE("/tags/:id", v1.DeleteTag)

		//获取文章列表
		apiV1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiV1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiV1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiV1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiV1.DELETE("/articles/:id", v1.DeleteArticle)
	}
	return r
}
