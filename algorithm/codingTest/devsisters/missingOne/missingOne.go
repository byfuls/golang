package main

func findMissingOne(a, b []int) int {
	result := 0
	for _, n := range a {
		result ^= n
	}
	for _, n := range b {
		result ^= n
	}
	return result
}

func main() {
	var (
		a = []int{44, 22, 3, 14, 55}
		b = []int{3, 14, 7, 22, 44, 55}
	)
	missing := findMissingOne(a, b)
	println("Missing number is:", missing) // Output: Missing number is: 6
}
