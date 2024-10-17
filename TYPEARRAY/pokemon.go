package main

import "fmt"

const max int = 100

type pokemon struct {
	nama  string
	jenis string
	level int
}

type ability struct {
	skill1 pokemon
}

func main() {
	var p [max]pokemon
	var a [max]ability
	/*fmt.Scan(&p.nama, &p.jenis, &p.level, &a.skill1.nama, &a.skill1.jenis, &a.skill1.level)
	fmt.Println(p.nama, p.jenis, p.level, a.skill1.nama, a.skill1.jenis, a.skill1.level)
	*/
	var i, n, indeks int
	fmt.Print("Jumlah Data Pokemon: ")
	fmt.Scan(&n)
	fmt.Println("Masukkan Data Pokemon")
	for i = 0; i < n; i++ {
		fmt.Print("Data Ke-", i, ": ")
		fmt.Scan(&p[i].nama, &p[i].jenis, &p[i].level, &a[i].skill1.nama, &a[i].skill1.jenis, &a[i].skill1.level)
	}
	fmt.Println("Isi Data")
	for i = 0; i < n; i++ {
		fmt.Print("Data Ke-", i, ": ")
		fmt.Println(p[i].nama, p[i].jenis, p[i].level, a[i].skill1.nama, a[i].skill1.jenis, a[i].skill1.level)
	}
	fmt.Print("Kode Indeks Data Pokemon: ")
	fmt.Scan(&indeks)
	fmt.Println(p[indeks].nama, p[indeks].jenis, p[indeks].level, a[indeks].skill1.nama, a[indeks].skill1.jenis, a[indeks].skill1.level)

}
