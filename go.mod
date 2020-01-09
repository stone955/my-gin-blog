module github.com/stone955/my-gin-blog

go 1.13

replace (
	github.com/stone955/my-gin-blog/do => ./router/api
	github.com/stone955/my-gin-blog/docs => ./docs
	github.com/stone955/my-gin-blog/docs/api => ./router/api
	github.com/stone955/my-gin-blog/middleware => ./middleware
	github.com/stone955/my-gin-blog/middleware/jwt => ./middleware/jwt
	github.com/stone955/my-gin-blog/model => ./model
	github.com/stone955/my-gin-blog/pkg/e => ./pkg/e
	github.com/stone955/my-gin-blog/pkg/setting => ./pkg/setting
	github.com/stone955/my-gin-blog/router => ./router
	github.com/stone955/my-gin-blog/router/api => ./router/api
)

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fvbock/endless v0.0.0-20170109170031-447134032cb6
	github.com/gin-gonic/gin v1.5.0
	github.com/go-ini/ini v1.51.1
	github.com/jinzhu/gorm v1.9.11
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/swaggo/swag v1.6.4
	github.com/unknwon/com v1.0.1
	gopkg.in/ini.v1 v1.51.0 // indirect
	gopkg.in/validator.v2 v2.0.0-20191107172027-c3144fdedc21
)
