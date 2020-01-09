package model

type Auth struct {
	ID       uint `gorm:"primary_key"`
	Username string
	Password string
}

func CheckAuth(username, password string) bool {
	var auth Auth
	db.Select("id").Where(Auth{
		Username: username,
		Password: password,
	}).First(&auth)

	if auth.ID > 0 {
		return true
	}
	return false
}
