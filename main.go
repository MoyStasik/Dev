package main

import "fmt"

func fibonacci(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Print(fibonacci(10))
}
