package basic

import (
	"fmt"
	"strconv"
)

type userInfo struct {
	id        int
	firstName string
	lastName  string
	email     string
	userRate  float64
	mobile    string
}

func getUserRate(userID int) float64 {
	if userID == 155 {
		return 499998.23
	}
	return 0.0

}

// (value receiver) Method
func (u userInfo) information() string {
	return "Hi! My name is " + u.firstName + " " + u.lastName + ". My UserID is " + strconv.Itoa(u.id) + ". My email address " + u.email + ". Hourly wage => " + strconv.FormatFloat(u.userRate, 'f', 6, 64) + "."
}

// (pointer receiver) Method
func (u *userInfo) updateEmail() {
	u.email = "matta@yoyocode.dev"
}

// (pointer receiver) Method to update userRate
func (u *userInfo) updateUserRate(userRate float64) {
	u.userRate = userRate
}

/*Structs sample userInfo*/
func Structs() {
	user1 := userInfo{id: 155, firstName: "Harish", lastName: "Matta", email: "harish@yoyocode.dev", mobile: "+91 9866863834"}
	user1.updateUserRate(getUserRate(user1.id))
	fmt.Println(user1)
	fmt.Println(user1.information())
}
