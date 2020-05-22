package models

type User struct {
	id int64
	Name string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}
