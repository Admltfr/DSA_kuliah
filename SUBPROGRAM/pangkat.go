package main

import "fmt"

func pangkat(m, n int) int {
	/*{Terdefinisi sebuah bilangan bulat positif m dan n.
	  Fungsi mengembalikan hasil m^n. }*/
	if n == 1 {
		return m
	} else {
		return m * pangkat(m, n-1)
	}

}

func main() {
	var n, m int
	fmt.Scan(&m, &n)
	fmt.Println(pangkat(m, n))
}
