package main

import "fmt"

func kali2(n int) int {
	if n == 1 {
		return 2
	} else {
		return 2*n + kali2(n-1)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(kali2(n))
}
