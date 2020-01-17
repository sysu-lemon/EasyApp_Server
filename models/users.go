package models

type User struct {
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}
