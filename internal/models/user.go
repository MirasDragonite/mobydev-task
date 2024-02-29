package models

type User struct {
	Id             int64  `json:"id"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	HashedPassword string `json:"password"`
	MobilePhone    string `json:"mobilePhone"`
	BirthDate      string `json:"createdDate"`
}

type Login struct {
	Email      string `json:"email"`
	Passowrd   string `json:"password"`
	Repassword string `json:"repeatPassword"`
}

type Register struct {
	Email    string `json:"email"`
	Passowrd string `json:"password"`
}

type Session struct {
	Id          int64
	UserId      int64
	Token       string
	ExpiredDate string
}
