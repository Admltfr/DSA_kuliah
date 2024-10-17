package main

import "fmt"

func modulusLoop(n, m int) int {
	for n >= m {
		n = n - m
	}
	return n
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	fmt.Println(modulusLoop(n, m))
}
