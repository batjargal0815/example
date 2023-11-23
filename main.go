package main

import "fmt"

type User struct {
	UserName  string `json:"user_name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {
	user := &User{
		UserName:  "batja",
		LastName:  "Batjargal",
		FirstName: "Chuluunkhuu",
	}

	fmt.Println("Userinfo:", getUserInfo(user))
}

func getUserInfo(user *User) string {
	userInfo := user.UserName + " " + user.FirstName + " " + user.LastName
	return userInfo
}
