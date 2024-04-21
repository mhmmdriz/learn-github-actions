package main

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	println(sum(numbers))
}

func sum(numbers []int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}
