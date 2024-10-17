package main

import "fmt"

func sumdivmod(n int) int {
	if n == 0 {
		return 0
	} else {
		return n%10 + sumdivmod(n/10)
	}
}
func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(sumdivmod(n))
}
