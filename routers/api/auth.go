package api

import (
	"ginDemo/models"
	"ginDemo/pkg/e"
	"ginDemo/pkg/logging"
	"ginDemo/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	userName := c.Query("username")
	passowrd := c.Query("password")

	validator := validation.Validation{}

	a := auth{Username: userName, Password: passowrd}

	ok, _ := validator.Valid(&a)

	data := make(map[string]interface{})

	code := e.SUCCESS

	if !ok {
		code = e.INVALID_PARAMS
		for _, err := range validator.Errors {
			logging.Info(err.Key, err.Message)
			//log.Info(err.Key, err.Message)
		}
	}
	checkAuth := models.CheckAuth(userName, passowrd)
	if !checkAuth {
		code = e.ERROR_AUTH
		logging.Info("CheckAuth fail, Username or password wrong")
	}

	token, err := util.GenerateToken(userName, passowrd)
	if err != nil {
		code = e.ERROR_AUTH_TOKEN
		logging.Info("Generate token fail : %v", err)
	}
	data["token"] = token

	logging.Info("Generate Token : %s", token)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}
