package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	cetak(n, m)
}

func cetak(n, m int) {
	/*
		I.S. terdefinisi dua buah bilangan bulat positif n dan m, dengan n < m
		F.S. menampilkan barisan bilangan dari n hingga m
	*/
	var i, temp int
	if n > m {
		temp = n
		n = m
		m = temp

	}
	for i = m; i >= n; i-- {
		fmt.Printf("%d ", i)
	}
}
