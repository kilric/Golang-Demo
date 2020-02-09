package util

type User struct {
	Id   int
	Name string
}

type Admin struct {
	User
	level string
}
