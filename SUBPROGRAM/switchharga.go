package main

import "fmt"

func belanja(x, y int) int {
	var jumlah int
	if y >= 1 {
		switch x {
		case 1:
			jumlah = y * 2980
		case 2:
			jumlah = y * 4500
		case 3:
			jumlah = y * 9980
		case 4:
			jumlah = y * 4490
		case 5:
			jumlah = y * 6870
		default:
			jumlah = 0
		}
	} else {
		jumlah = 0
	}
	return jumlah

}
func main() {
	var noProduk, banyakBelanja, total int
	fmt.Scan(&noProduk, &banyakBelanja)
	total = belanja(noProduk, banyakBelanja)
	fmt.Println(total)
}
