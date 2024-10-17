package main

import "fmt"

func main() {
	var angkaDitebak int
	fmt.Scan(&angkaDitebak)
	tebakAngka(angkaDitebak)
}

func tebakAngka(angkaDitebak int) {
	/* I.S. terdefinisi angkaDitebak berisi bilangan bulat yang akan ditebak
	   F.S. program mengeluarkan banyak percobaan menebak angka hingga berhasil */
	var i, n int
	fmt.Scan(&n)
	i = 1
	for angkaDitebak != n && angkaDitebak != n%10 && angkaDitebak != n/10 {
		fmt.Scan(&n)
		i++
	}
	fmt.Println(i)
}
