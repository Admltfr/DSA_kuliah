package main

import "fmt"

func cetakterbalikdigit(n int) {
	/* Terdefinisi sebuah bilangan bulat n.
	Tercetak tiap digit bilangan n diselingi dengan karakter dash (strip). */
	// lakukan sesuai dengan pseudocode
	if n < 10 {
		fmt.Print(n, "-")
	} else {
		fmt.Print(n%10, "-")
		cetakterbalikdigit(n / 10)
	}
}
func main() {
	var n int
	fmt.Scan(&n)
	cetakterbalikdigit(n)
}
