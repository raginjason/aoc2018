package main

// Sum variadic sum function
func Sum(numbers ...int) int {
	var total int
	for _, number := range numbers {
		total = total + number
	}
	return total
}
