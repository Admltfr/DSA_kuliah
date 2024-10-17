package main

import "fmt"

func gunung(x, y, n int) int {

	//fungsi akan mengembalikan jumlah letusan
	//return merupakan berupa bilangan bulat yang menyatakan jumlah letusan.
	var i, jumlah int
	i = x
	for i <= n {
		jumlah = jumlah + 1
		i = i + x + y
	}
	i = y
	for i <= n {
		jumlah = jumlah + 1
		i = i + x + y
	}

	return jumlah

}

func main() {
	var x, y, n int
	fmt.Scan(&x, &y, &n)
	fmt.Println(gunung(x, y, n))
}
