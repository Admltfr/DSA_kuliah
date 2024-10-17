package main

import "fmt"

const NMAX int = 20

type tabChar [NMAX]byte

func main() {
	var kata tabChar
	var nChar int
	fmt.Scan(&nChar)
	if nChar > NMAX {
		nChar = NMAX
	}
	baca(&kata, &nChar)
	cetak(kata, nChar)
}

func baca(k *tabChar, n *int) {
	/*IS: Array karakter (k) dan banyak elemen (n) terdefinisi sembarang
	FS: Array karakter (k) dan banyak elemen (n) terdefinisi
	*/
	var i int
	var input byte
	for i = 0; i < *n; i++ {
		fmt.Scanf("%c", &input)
		(*k)[i] = input
	}
}

func cetak(k tabChar, n int) {
	/*
		IS: Array k dan n terdefinisi
		FS: Array k tercetak di layar dengan urutan terbalik
	*/
	var i int
	for i = n - 1; i >= 0; i-- {
		fmt.Printf("%c", k[i])
	}
}
