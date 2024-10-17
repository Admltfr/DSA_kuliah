package main

import "fmt"

type Mahasiswa struct {
	Nama string
	NIM  int
}

type tabMhs [40]Mahasiswa

func main() {
	var daftarMhs tabMhs
	var i int
	//  lakukan perulangan untuk menginputkan nama mahasiswa dan nim kemudian memasukkannya ke dalam array //
	i = 0
	fmt.Scan(&daftarMhs[i].Nama, &daftarMhs[i].NIM)
	for i < 40 {
		i++
		fmt.Scan(&daftarMhs[i].Nama)
		if daftarMhs[i].Nama == "#" {
			break
		}
		fmt.Scan(&daftarMhs[i].NIM)

	}
	var j int
	var s string
	fmt.Scan(&s)
	if s == "genap" {
		for j = 0; j < i; j++ {
			if daftarMhs[j].NIM%2 == 0 {
				fmt.Print(daftarMhs[j].Nama, " ")
			}

		}
	} else if s == "ganjil" {
		for j = 0; j < i; j++ {
			if daftarMhs[j].NIM%2 == 1 {
				fmt.Print(daftarMhs[j].Nama, " ")
			}

		}
	}

}
