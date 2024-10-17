package main

import "fmt"

func deretm(x int) int {
	if x == 1 {
		return 1
	} else if x > 1 && x%2 == 1 {
		return (-1 * x) + deretm(x-1)
	} else {
		return x + deretm(x-1)
	}
}

func main() {
	var x int
	fmt.Scan(&x)
	fmt.Println(deretm(x))
}
