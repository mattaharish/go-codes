package basic

import (
	"fmt"
	"go-codes/basic/numutil"
)

/*Array Logs array related functions*/
func Array() {
	// Arrays
	a := []float32{1, 2, 3, 4, 5, 6, 7, 8}
	b := []int{1, 2, 3, 4, 5}

	var avg float32
	avg = numutil.AvgOfArray(a)
	fmt.Printf("Sum of int array: %v\n", numutil.SumOfArray(b))
	fmt.Printf("Average of float array: %v\n", avg)
}
