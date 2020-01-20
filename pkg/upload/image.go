package upload

import (
	"fmt"
	"github.com/stone955/my-gin-blog/pkg/file"
	"github.com/stone955/my-gin-blog/pkg/logging"
	"github.com/stone955/my-gin-blog/pkg/setting"
	"github.com/stone955/my-gin-blog/pkg/util"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

// GetImageFullUrl 获取图片完整访问URL
func GetImageFullUrl(name string) string {
	return filepath.Join(setting.AppCfg.ImagePrefixUrl, GetImagePath(), name)
}

// GetImageName 获取图片名称
func GetImageName(name string) string {
	ext := filepath.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)
	return fileName + ext
}

// GetImagePath 获取图片路径
func GetImagePath() string {
	return setting.AppCfg.ImageSavePath
}

// GetImageFullPath 获取图片完整路径
func GetImageFullPath() string {
	return filepath.Join(setting.AppCfg.RuntimeRootPath, GetImagePath())
}

// CheckImageExt 检查图片后缀
func CheckImageExt(fileName string) bool {
	ext := filepath.Ext(fileName)
	for _, allowExt := range setting.AppCfg.ImageAllowExtNames {
		if strings.ToUpper(ext) == strings.ToUpper(allowExt) {
			return true
		}
	}
	return false
}

// CheckImageSize 检查图片大小
func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Println(err)
		logging.Warn(err)
		return false
	}

	return size <= setting.AppCfg.ImageMaxSize
}

// CheckImage 检查图片
func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = file.MkDirIfNotExist(filepath.Join(dir, src))
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
