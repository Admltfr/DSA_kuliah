package main

import "fmt"

func sama(n, m int) {
	fmt.Scan(&n, &m)
	for n != m {
		fmt.Println(m, n)
		fmt.Scan(&n, &m)
	}
}

func main() {
	var n, m int
	sama(n, m)
}
