package main

import "fmt"

func barisganjil(x, angka int) int {
	if x <= angka {
		if x%2 != 1 {
			return barisganjil(x+1, angka)
		} else if x%2 == 1 {
			return x + barisganjil(x+1, angka)
		}
	}
	return 0
}

func main() {
	var x, angka int
	fmt.Scan(&angka)
	fmt.Println((barisganjil(x, angka)))

}
