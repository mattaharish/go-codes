package basic

import (
	"fmt"
)

/*Maps List of emails*/
func Maps() {
	// Maps
	userDetails := make(map[int]string)

	userDetails[155] = "harish@yoyocode.dev"
	userDetails[156] = "matta@yoyocode.dev"
	userDetails[157] = "kumar@yoyocode.dev"

	fmt.Println(userDetails)

	for index, data := range userDetails {
		fmt.Printf("%d -> %s\n", index, data)
	}
}
