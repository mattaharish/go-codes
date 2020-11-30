package basic

import (
	"fmt"
	"go-codes/basic/numutil"
	"math"
)

func square(n int) float64 {
	return math.Sqrt(25)
}

/*MathOperators Math operations */
func MathOperators() {
	n := 10
	fmt.Printf("%T", n)
	fmt.Println(square(n))
	fmt.Println(numutil.Prime(11))
	fmt.Println(numutil.Odd(10))
	fmt.Println(numutil.Even(10))
}
