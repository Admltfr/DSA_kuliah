package main

import "fmt"

func lowercase(s string, n int) {
	if n > 0 {
		lowercase(s, n-1)
		fmt.Print(string(s[n-1] + 32))
	}
}

func main() {
	var s string
	fmt.Scan(&s)
	lowercase(s, len(s))
}
