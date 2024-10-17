package main

import "fmt"

func dectobin(n int) {
	/* Terdefinisi sebuah bilangan bulat positif n.
	Tercetak bentuk biner dari bilangan n. */
	if n/2 == 0 {
		fmt.Print(n)
	} else {
		dectobin(n / 2)
		fmt.Print(n % 2)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	dectobin(n)
}
