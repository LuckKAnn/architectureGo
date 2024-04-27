package models

type User struct {
	Id       int64
	Username string
	Password string
	Email    string
	Age      int32
	Sex      string
	Tel      string
	Addr     string
	Card     string
	Married  int
	Salary   float64
}

func (user User) TableName() string {
	return "user"
}
