package main

import (
	"fmt"
)

type User struct {
	Name, Surname   string
	Age             int
	Email, Password string
}

func (u *User) ChangeName(name string) {
	u.Name = name
}
func (u *User) ChangeSurname(surname string) {
	u.Surname = surname
}
func (u *User) ChangeAge(age int) {
	u.Age = age
}
func (u *User) ChangeEmail(email string) {
	u.Email = email
}
func (u *User) ChangePassword(password string) {
	u.Password = password
}

func main() {
	user := User{
		Name:     "Pepe",
		Surname:  "Perez",
		Age:      20,
		Email:    "pepe@example.com",
		Password: "123456789",
	}
	user.ChangeName("user")
	user.ChangeSurname("user")
	user.ChangeAge(22)
	user.ChangeEmail("user@example.com")
	user.ChangePassword("password")

	fmt.Println(user)
}
