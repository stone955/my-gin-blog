package api

import (
	"github.com/gin-gonic/gin"
	"github.com/stone955/my-gin-blog/pkg/e"
	"github.com/stone955/my-gin-blog/pkg/logging"
	"github.com/stone955/my-gin-blog/pkg/upload"
	"net/http"
	"path/filepath"
)

func UploadImage(c *gin.Context) {
	code := e.Ok
	data := make(map[string]string)

	file, image, err := c.Request.FormFile("image")
	if err != nil {
		logging.Warn(err)
		code = e.Error
		c.JSON(http.StatusOK, H(code, data))
		return
	}

	if image == nil {
		code = e.InvalidParams
	} else {
		imageName := upload.GetImageName(image.Filename)
		fullPath := upload.GetImageFullPath()
		savePath := upload.GetImagePath()

		if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
			code = e.ErrorUploadCheckImageFormat
		} else {
			src := filepath.Join(fullPath, imageName)
			if err := upload.CheckImage(src); err != nil {
				logging.Warn(err)
				code = e.ErrorUploadCheckImageFail
			} else if err := c.SaveUploadedFile(image, savePath); err != nil {
				logging.Warn(err)
				code = e.ErrorUploadSaveImageFail
			} else {
				data["image_url"] = upload.GetImageFullUrl(imageName)
				data["image_save_url"] = filepath.Join(savePath, imageName)
			}
		}
	}

	c.JSON(http.StatusOK, H(code, data))
}
