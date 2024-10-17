package main

import "fmt"

func jumlahan(x, n int, jumlah *int) {
	if x > 0 {
		fmt.Scan(&n)
		*jumlah = *jumlah + n
		jumlahan(x-1, n, jumlah)
	}
}

func main() {
	var x, n, jumlah int
	fmt.Scan(&x)
	jumlahan(x, n, &jumlah)
	fmt.Println(jumlah)
	if jumlah%2 == 0 {
		fmt.Println(true)
	} else if jumlah%2 == 1 {
		fmt.Println(false)
	}
}
