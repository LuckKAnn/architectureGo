package api

import (
	"ginDemo/pkg/e"
	"ginDemo/pkg/logging"
	"ginDemo/pkg/upload"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadImage(c *gin.Context) {

	code := e.SUCCESS
	data := make(map[string]interface{})
	file, image, err := c.Request.FormFile("image")
	if err != nil {
		code = e.ERROR
		logging.Warn(err)
	}

	if image == nil {
		code = e.INVALID_PARAMS
	} else {
		// 各种去检查文件的大小和一些信息
		imageName := upload.GetImageName(image.Filename)
		fullPath := upload.GetImageFullPath()
		savePath := upload.GetImagePath()
		src := fullPath + imageName

		if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
			// 文件格式和文件大小不符合要求
			code = e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT
		} else {
			// 这里是检查存储目录是否存在，是否具有写入的权限
			err := upload.CheckImage(fullPath)
			if err != nil {
				logging.Warn(err)
				code = e.ERROR_UPLOAD_CHECK_IMAGE_FAIL
			} else if err := c.SaveUploadedFile(image, src); err != nil {
				logging.Warn(err)
				code = e.ERROR_UPLOAD_SAVE_IMAGE_FAIL
			} else {
				data["image_url"] = upload.GetImageFullUrl(imageName)
				data["image_save_url"] = savePath + imageName
			}
		}

	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}
