package main

import "fmt"

func cekGenap(n int) {
	/*
		I.S. terdefinisi bilangan bulat positif.
		F.S. tercetak "Ya" bila bilangan masukan adalah genap atau "Tidak" bila bukan.
	*/
	if n > 0 {
		if n%2 == 0 {
			fmt.Println("Ya")
		} else if n%2 == 1 {
			fmt.Println("Tidak")
		}
	}
}

func main() {
	var bilangan int
	fmt.Scan(&bilangan)
	cekGenap(bilangan)
}
