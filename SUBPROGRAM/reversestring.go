package main

import "fmt"

func reverse(str string, n int) {
	/*{ I.S. Terdefinisi str sebagai input string
	F.S. menampilkan string yang terbalik menggunakan fungsi rekursif*/
	if n > 0 {
		fmt.Print(string(str[n-1]))
		reverse(str, n-1) // hint :  gunakan var temporary
	}
}
