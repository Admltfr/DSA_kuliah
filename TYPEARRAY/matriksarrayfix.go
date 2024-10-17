package main

import "fmt"

const N = 100

type kolom [N]int

func main() {
	var baris [N]kolom
	var x int
	var vertikal, horizontal, matriks, dimensi int
	fmt.Scan(&dimensi)
	for vertikal = 0; vertikal < dimensi; vertikal++ {
		for horizontal = 0; horizontal < dimensi; horizontal++ {
			fmt.Scan(&matriks)
			baris[vertikal][horizontal] = matriks
		}
	}
	fmt.Scan(&x)
	searchMatrix(&x, baris)
}

func searchMatrix(x *int, baris [N]kolom) {
	var vertikal, horizontal, count, hasil int
	for vertikal = 0; vertikal < N; vertikal++ {
		for horizontal = 0; horizontal < N; horizontal++ {
			hasil = baris[vertikal][horizontal]
			if hasil == *x && count == 0 {
				fmt.Println(vertikal, horizontal)
				count++
			}
		}
	}
}
