package main

import "fmt"

func string1(s string, n int) {
	if n == 0 {
		fmt.Printf("%c", s[0])
	} else {
		fmt.Printf("%c", s[n])
		string1(s, n-1)
	}
}
func main() {
	var n int
	var s string
	fmt.Scan(&s)
	n = len(s) - 1
	string1(s, n)
}
