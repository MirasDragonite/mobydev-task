package models

type User struct {
	Id             int64  `json:"id"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	HashedPassword string `json:"password"`
	BirthDate      string `json:"createdDate"`
	MobilePhone    string `json:"mobilePhone"`
}
