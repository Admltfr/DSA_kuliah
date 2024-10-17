package main

import "fmt"

func hitungSKS(totalSKS int) int {
	var kodeMatkul string
	var sks int
	fmt.Scan(&kodeMatkul)
	if kodeMatkul == "0" {
		return totalSKS
	} else {
		sks = 0
		if kodeMatkul[0] == 'A' || kodeMatkul[0] == 'B' || kodeMatkul[0] == 'C' {
			sks = 2
		} else if kodeMatkul[0] == 'D' || kodeMatkul[0] == 'E' || kodeMatkul[0] == 'F' {
			sks = 3
		}

		return hitungSKS(totalSKS + sks)
	}
}

func main() {
	var totalSKS, sks int
	sks = 0
	totalSKS = hitungSKS(sks)
	fmt.Print(totalSKS)
}
