package e

var msgFlags = map[int]string{
	Create:                     "created",
	Ok:                         "ok",
	Delete:                     "deleted",
	Error:                      "error",
	InvalidParams:              "invalid params",
	ErrorExistTag:              "tag exist",
	ErrorNotExistTag:           "tag not exist",
	ErrorNotExistArticle:       "article not exist",
	ErrorAuthCheckTokenFail:    "token check fail",
	ErrorAuthCheckTokenTimeout: "token expired",
	ErrorAuthToken:             "token auth fail",
	ErrorAuth:                  "token error",
}

func GetMsg(code int) string {
	if msg, ok := msgFlags[code]; ok {
		return msg
	}
	return msgFlags[Error]
}
