package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthday  string
	createAt  time.Time
}

func (u user) outputUserData() {
	fmt.Println(u.firstName, u.lastName, u.birthday)
}
func (u *user) deleteUserName() {
	u.firstName = ""
	u.lastName = ""
}

func main() {
	var userFirstName string
	var userLastName string
	var userBirthday string

	getDataFromUser(userFirstName)
	getDataFromUser(userLastName)
	getDataFromUser(userBirthday)

	var appData = user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthday:  userBirthday,
		createAt:  time.Now(),
	}

	appData.outputUserData()
	appData.deleteUserName()
	appData.outputUserData()
}
func getDataFromUser(text string) {
	fmt.Println(":")
	fmt.Scan(&text)
}
