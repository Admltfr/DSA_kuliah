package main

import "fmt"

func main() {
	var bil1, bil2 int
	fmt.Scan(&bil1, &bil2)
	fmt.Println(jumlahGenap(bil1, bil2))
}

func jumlahGenap(n, m int) int {
	/* mengembalikan penjumlahan bilangan ganjil dalam interval bil1 dan bil2 */
	var i, hasil int
	for i = n; i <= m; i++ {
		if i%2 == 0 {
			hasil = hasil + i
		}
	}
	return hasil
}
