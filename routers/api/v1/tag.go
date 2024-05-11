package v1

import (
	"ginDemo/models"
	"ginDemo/pkg/e"
	"ginDemo/pkg/setting"
	"ginDemo/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

// 获取多个文章标签
func GetTags(c *gin.Context) {
	name := c.Query("name")
	data := make(map[string]interface{})
	maps := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}
	code := e.SUCCESS

	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// 新增文章标签
func AddTag(c *gin.Context) {
	// 获取参数
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("create_by")

	// 参数校验
	// 居然现在还是手动写参数校验
	validator := validation.Validation{}
	validator.Required(name, "name").Message("名称不能为空")
	validator.MaxSize(name, 100, "name").Message("名称最长为100字符")
	validator.Required(createdBy, "created_by").Message("创建人不能为空")
	validator.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	validator.Range(state, 0, 1, "state").Message("状态只允许0或1")
	// 设置代码状态
	// 如果没有tag，创建
	code := e.INVALID_PARAMS
	if !validator.HasErrors() {
		if !models.ExistTagByName(name) {
			code = e.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// 修改文章标签
func EditTag(c *gin.Context) {
}

// 删除文章标签
func DeleteTag(c *gin.Context) {
}