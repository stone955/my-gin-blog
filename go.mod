module github.com/stone955/my-gin-blog

go 1.13

replace (
	github.com/stone955/my-gin-blog/middleware => ./middleware
	github.com/stone955/my-gin-blog/model => ./model
	github.com/stone955/my-gin-blog/pkg/e => ./pkg/e
	github.com/stone955/my-gin-blog/pkg/setting => ./pkg/setting
	github.com/stone955/my-gin-blog/router => ./router
)

require (
	github.com/gin-gonic/gin v1.5.0
	github.com/go-ini/ini v1.51.1
	github.com/jinzhu/gorm v1.9.11
	github.com/kr/pretty v0.1.0 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/unknwon/com v1.0.1
	gopkg.in/ini.v1 v1.51.0 // indirect
	gopkg.in/validator.v2 v2.0.0-20191107172027-c3144fdedc21
)
