package e

const (
	Create = 201
	Ok     = 200
	Delete = 204

	InvalidParams = 400
	Error         = 500

	ErrorExistTag        = 10001
	ErrorNotExistTag     = 10002
	ErrorNotExistArticle = 10003

	ErrorAuthCheckTokenFail    = 20001
	ErrorAuthCheckTokenTimeout = 20002
	ErrorAuthToken             = 20003
	ErrorAuth                  = 20004
)
