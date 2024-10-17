package main

import "fmt"

const N = 100

type kolom [N]int

func main() {
	var barisan [N]kolom
	var x int
	fmt.Scan(&x)
	searchMatrix(&x, barisan)
}

func searchMatrix(x *int, baris [N]kolom) {
	var i, j, matriks, dimensi, hasil, count int
	fmt.Scan(&dimensi)
	for i = 0; i < dimensi; i++ {
		for j = 0; j < dimensi; j++ {
			fmt.Scan(&matriks)
			baris[i][j] = matriks
			fmt.Println(baris[i][j])
		}
	}
	i = 0
	j = 0

	for i = 0; i < dimensi; i++ {
		for j = 0; j < dimensi; j++ {
			hasil = baris[i][j]
			if hasil == *x && count == 0 {
				fmt.Println(i, j)
				count++
			}
		}
	}
}
