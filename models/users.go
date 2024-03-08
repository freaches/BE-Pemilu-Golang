package models

type User struct {
	ID       uint
	Name     string
	Address  string
	Gender   string
	Username string
	Password string
	Role     string
}

type UserArticle struct {
	ID   uint
	Name string
}

type UserPaslon struct {
	ID      uint
	Name    string
	Address string
	Gender  string
}

func (UserArticle) TableName() string {
	return "users"
}

func (UserPaslon) TableName() string {
	return "users"
}
