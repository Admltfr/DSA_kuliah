package main

import "fmt"

func sum1(n int) int {
	if n == 1 {
		return n
	} else {
		return n + sum1(n-1)
	}
}
func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(sum1(n))
}
