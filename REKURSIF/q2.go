package main

import "fmt"

func print2(n int) {
	if n > 0 {
		fmt.Println(n)
		print2(n - 1)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	print2(n)
}
