package main

import (
	"fmt"
	"math"
)

func main() {
	var n, m int
	bilanganPola(n, m)
}
func bilanganPola(n, m int) {
	/*
	   IS : Terdefinisi bilangan bulat n dan m
	   FS : menampilkan pola barisan bilangan
	*/
	fmt.Scan(&n, &m)
	var i int
	for i = n; i <= m; i++ {
		if i%2 != 0 {
			fmt.Printf("%.2f %s", math.Pow(float64(2), float64(i)), " ")
		} else if i%2 == 0 {
			fmt.Printf("%.2f %s", math.Pow(float64(2), (-1*float64(i))), " ")
		}
	}

}
