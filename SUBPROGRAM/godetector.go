package main

import "fmt"

func cariGo(kar byte) {
	/* I.S. terdefinisi kar berisi serangkaian karakter
	   F.S. program mengeluarkan jumlah karakter go yang ada dalam karakter */
	var kar2 byte
	var jumlah int
	for kar != '.' {
		if kar2 == 'g' && kar == 'o' {
			jumlah = jumlah + 1
		}
		kar2 = kar
		fmt.Scanf("%c", &kar)
	}
	fmt.Println(jumlah)
}

func main() {
	var kar byte
	fmt.Scanf("%c", &kar)
	cariGo(kar)
}
