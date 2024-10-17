package main

import "fmt"

func totalgenap(n int) int {
	if n == 0 {
		return n
	} else if n%2 != 0 {
		return totalgenap(n - 1)
	}
	return n + totalgenap(n-1)
}
func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(totalgenap(n))

}
