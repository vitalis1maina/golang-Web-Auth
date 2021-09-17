package main

import (
	"fmt"
)

func main() {
	n := 5
	for i := 0; i <= n; i++ {
		fmt.Println(fibonacci(i))
	}
}
func fibonacci(n int) int {
	if n == 0 || n == 1 {
		return 0
	} else if n == 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2) + fibonacci(n-3)
}
