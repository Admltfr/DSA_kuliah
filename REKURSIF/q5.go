package main

import "fmt"

func binary1(n int) int {
	if n == 0 {
		return 0
	} else {
		return n%2 + 10*binary1(n/2)
	}
}
func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(binary1(n))
}
