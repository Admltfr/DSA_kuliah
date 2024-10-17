package main

import "fmt"

type pembeli struct {
	member bool
	harga  int
}

func main() {
	var p pembeli
	fmt.Scan(&p.member, &p.harga)
	hitungBiaya(&p)
}
func hitungBiaya(p *pembeli) {
	var total float64
	if p.harga >= 100000 && p.harga < 300000 {
		total = float64(p.harga) - (float64(p.harga) * 0.1)
	} else if p.harga >= 300000 {
		total = float64(p.harga) - (float64(p.harga) * 0.1)
	}

	if p.harga >= 200000 && p.harga < 400000 && p.member {
		total = total - 25000
	} else if p.harga >= 400000 && p.member {
		total = total - 50000
	}
	fmt.Println(total)
}
