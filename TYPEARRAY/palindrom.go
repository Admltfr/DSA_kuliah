package main

import "fmt"

const N int = 100

type tab [N]rune

func masukanArray(T *tab, n *int) {
	/*I.S. Data tersedia dalam piranti masukan
	  F.S. Array t berisi sejumlah n karakter yang dimasukka user,
	  Proses input selama karakter bukanlah TITIK dan n <= N
	  Petunjuk: Gunakan fmt.Scanf untuk melakukan input karakter
	*/
	var input rune
	*n = 0
	fmt.Scanf("%c", &input)
	for input != '.' && *n <= N {
		(T)[*n] = input
		*n++
		fmt.Scanf("%c", &input)
	}

}

func reverseArray(T1 tab, n int, T2 *tab) {
	/*I.S. Terdefinisi array T yang berisikan n jumlah karakter
	  F.S. Isi array sudah direverse (urutan isi array T terbalik)
	*/
	var i, j int
	j = n - 1
	for i = 0; i < n; i++ {
		T2[i] = T1[j]
		j--
	}

}

func isPalindrom(T tab, n int) bool {
	/*Mengembalikan true apabila susunan karakter di dalam t
	  membentuk palindromnya, dan false apabila sebaliknya
	  Petunjuk: Manfaatkan prosesdur reverseArray*/
	var temp tab
	var i int = 0
	var status bool = true
	reverseArray(T, n, &temp)
	for i < n && status {
		status = T[i] == temp[i]
		i++
	}
	return status
}

func main() {
	var T tab
	var n int
	masukanArray(&T, &n)
	fmt.Println(isPalindrom(T, n))
}
