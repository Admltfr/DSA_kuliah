package main

import "fmt"

func print1(n int) {
	if n > 0 {
		print1(n - 1)
		fmt.Println(n)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	print1(n)
}
