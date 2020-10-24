package array

// Sum1 gets the ans of arr with common for loop
func Sum1(arr [5]int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += arr[i]
	}
	return sum
}

// Sum2 gets the ans of arr with for range loop
func Sum2(arr [5]int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}
