package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}

func main() {
	res := fact(11)
	fmt.Println("fact(11) = ", res)

	res = fact(9)
	fmt.Println("fact(9) = ", res)
}
