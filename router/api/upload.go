package api

import (
	"github.com/gin-gonic/gin"
	"github.com/stone955/my-gin-blog/pkg/app"
	"github.com/stone955/my-gin-blog/pkg/e"
	"github.com/stone955/my-gin-blog/pkg/logging"
	"github.com/stone955/my-gin-blog/pkg/upload"
	"net/http"
	"path/filepath"
)

func UploadImage(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		data = make(map[string]string)
	)

	file, image, err := c.Request.FormFile("image")
	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.Error, data)
		return
	}

	if image == nil {
		appG.Response(http.StatusBadRequest, e.InvalidParams, data)
		return
	}

	imageName := upload.GetImageName(image.Filename)
	fullPath := upload.GetImageFullPath()
	savePath := upload.GetImagePath()

	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		appG.Response(http.StatusBadRequest, e.ErrorUploadCheckImageFormat, data)
		return
	}

	src := filepath.Join(fullPath, imageName)
	if err := upload.CheckImage(src); err != nil {
		logging.Warn(err)
		appG.Response(http.StatusBadRequest, e.ErrorUploadCheckImageFail, data)
		return
	}
	if err := c.SaveUploadedFile(image, savePath); err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ErrorUploadSaveImageFail, data)
		return
	}

	data["image_url"] = upload.GetImageFullUrl(imageName)
	data["image_save_url"] = filepath.Join(savePath, imageName)
	appG.Response(http.StatusOK, e.Ok, data)
}
