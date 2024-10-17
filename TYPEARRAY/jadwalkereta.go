package main

import "fmt"

type data struct {
	jam    int
	menit  int
	nomor  string
	asal   string
	tujuan string
}

type kereta struct {
	k1, k2, k3 data
}

func main() {
	var m, k, n kereta
	fmt.Scan(&n.k1.nomor, &n.k1.asal, &m.k1.jam, &m.k1.menit, &n.k1.tujuan, &k.k1.jam, &k.k1.menit)
	fmt.Scan(&n.k2.nomor, &n.k2.asal, &m.k2.jam, &m.k2.menit, &n.k2.tujuan, &k.k2.jam, &k.k2.menit)
	fmt.Scan(&n.k3.nomor, &n.k3.asal, &m.k3.jam, &m.k3.menit, &n.k3.tujuan, &k.k3.jam, &k.k3.menit)
	m.k1.menit = m.k1.jam*60 + m.k1.menit
	m.k2.menit = m.k2.jam*60 + m.k2.menit
	m.k3.menit = m.k3.jam*60 + m.k3.menit
	k.k1.menit = k.k1.jam*60 + k.k1.menit
	k.k2.menit = k.k2.jam*60 + k.k2.menit
	k.k3.menit = k.k3.jam*60 + k.k3.menit
	n.k1.menit = k.k1.menit - m.k1.menit
	n.k2.menit = k.k2.menit - m.k2.menit
	n.k3.menit = k.k3.menit - m.k3.menit
	if n.k1.menit >= n.k2.menit && n.k1.menit >= n.k3.menit {
		fmt.Println(n.k1.nomor)
	} else if n.k2.menit >= n.k1.menit && n.k2.menit >= n.k3.menit {
		fmt.Println(n.k2.nomor)
	} else if n.k3.menit >= n.k1.menit && n.k3.menit >= n.k2.menit {
		fmt.Println(n.k3.nomor)
	}

}
