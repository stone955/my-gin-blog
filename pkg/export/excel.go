package export

import (
	"github.com/stone955/my-gin-blog/pkg/setting"
	"path"
)

func GetExcelFullUrl(name string) string {
	return path.Join(setting.AppCfg.PrefixUrl, GetExcelPath(), name)
}

func GetExcelPath() string {
	return setting.AppCfg.ExportSavePath
}

func GetExcelFullPath() string {
	return path.Join(setting.AppCfg.RuntimeRootPath, GetExcelPath())
}
