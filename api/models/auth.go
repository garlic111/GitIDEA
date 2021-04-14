package models

type User struct {
	Id       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Head   string `json:"head"`
	Address string `json:"address"`
}

func GetUserID(username string) string {
	var user User
	db.Select("username").Where(User{Username: username}).First(&user)
	return user.Username
}