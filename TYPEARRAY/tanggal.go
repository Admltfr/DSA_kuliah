package main

import "fmt"

func main() {
	type waktu struct {
		hari  int
		bulan int
		tahun int
	}

	var nk, nm, m, k waktu
	var jumlah int
	fmt.Scanf("%d %d %d \n %d %d %d", &nm.hari, &nm.bulan, &nm.tahun, &nk.hari, &nk.bulan, &nk.tahun)
	m.hari = nm.hari + nm.bulan*30 + nm.tahun*360
	k.hari = nk.hari + nk.bulan*30 + nk.tahun*360
	jumlah = k.hari - m.hari
	fmt.Println(jumlah)

}
