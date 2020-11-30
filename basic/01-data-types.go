package basic

import (
	"fmt"
	"go-codes/basic/numutil"
	"strconv"
)

/* basic data types */
func basicDataTypes() {
	var task = "Learn Go!"
	var days = 45
	var revision = true
	var duration = 30.5
	fmt.Println("Task:", task)
	fmt.Println("Days:", days)
	fmt.Println("Total Hours:", duration)
	fmt.Println("Revision:", revision)
}

func getCoolStatement(name string, age int, isCool bool, email string) string {
	return "Chill! This is " + name + ", aged " + strconv.Itoa(age) + " and I feel like, I'm cool " + strconv.FormatBool(isCool) + " and my email id " + email
}

/* swap two number */
func swap(x, y int) (int, int) {
	return y, x
}

/*Types Function that assigns and logs the basic types*/
func Types() {
	/* Log some basic data types  */
	basicDataTypes()

	name := "Matta"
	age := 24
	isCool := true
	email := "harish.matta31@gmail.com"

	/* Log type of variables */
	fmt.Printf("Type of name -> %T\n", name)
	fmt.Printf("Type of age -> %T\n", age)
	fmt.Printf("Type of isCool -> %T\n", isCool)
	fmt.Printf("Type of email -> %T\n", email)

	/* Get a cool statement */
	fmt.Println(getCoolStatement(name, age, isCool, email))

	/* Check if a number is prime or not */
	num := 2
	fmt.Printf("is %d a prime number? %v \n", num, numutil.Prime(num))

	/* swap two number */
	fmt.Println(swap(1, 2))
}
