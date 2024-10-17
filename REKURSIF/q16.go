package main

import "fmt"

func pangkat(m, n int) int {
	if n == 0 {
		return 1
	} else {
		return m * pangkat(m, n-1)
	}
}

func main() {
	var m, n int
	fmt.Scan(&m, &n)
	fmt.Println(pangkat(m, n))
}
