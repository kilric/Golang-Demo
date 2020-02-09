package util

import "fmt"

type User struct {
	Name  string
	Email string
}

type Notify interface {
	Notify()
}

type Out interface {
	Out()
}

func (user *User) Notify() {
	fmt.Println("name :", user.Name, "email:", user.Email)
}

func (user User) Out() {
	fmt.Println("name:", user.Name, "email:", user.Email)
}
