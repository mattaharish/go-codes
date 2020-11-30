package numutil

// Prime check if a given number  is prime or not
func Prime(num int) bool {
	if (num == 0) || (num == 1) {
		return false
	}
	counter := 0
	for index := 2; index < num/2; index++ {
		if num%index == 0 {
			counter++
		}
	}
	if counter > 0 {
		return false
	}
	return true
}

// Odd check if a given number is odd
func Odd(n int) bool {
	if n%2 == 0 {
		return false
	}
	return true

}

// Even check if a given number is even
func Even(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false

}

// SumOfArray get the sum of an array
func SumOfArray(arr []int) int {
	sum := 0
	for index := 0; index < len(arr); index++ {
		sum = sum + arr[index]
	}
	return sum
}

// AvgOfArray get the sum of an array
func AvgOfArray(arr []float32) float32 {
	var sum float32 = 0.0

	for _, num := range arr {
		sum += num
	}
	return sum / float32(len(arr))
}
