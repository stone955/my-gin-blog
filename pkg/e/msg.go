package e

var msgFlags = map[int]string{
	Success:                    "ok",
	Error:                      "error",
	InvalidParams:              "请求参数错误",
	ErrorExistTag:              "已存在该标签名称",
	ErrorNotExistTag:           "该标签不存在",
	ErrorNotExistArticle:       "该文章不存在",
	ErrorAuthCheckTokenFail:    "Token鉴权失败",
	ErrorAuthCheckTokenTimeout: "Token已超时",
	ErrorAuthToken:             "Token生成失败",
	ErrorAuth:                  "Token错误",
}

func GetMsg(code int) string {
	if msg, ok := msgFlags[code]; ok {
		return msg
	}
	return msgFlags[Error]
}
