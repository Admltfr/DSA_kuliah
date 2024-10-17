package main

import "fmt"

type Waktu struct {
	jam   int
	menit int
	detik int
}

func main() {
	var n, m, k Waktu
	fmt.Scan(&m.jam, &m.menit, &m.detik)
	fmt.Scan(&k.jam, &k.menit, &k.detik)
	m.detik = m.detik + (m.menit * 60) + (m.jam * 3600)
	k.detik = k.detik + (k.menit * 60) + (k.jam * 3600)
	n.detik = k.detik - m.detik

	n.jam = n.detik / 3600
	n.detik = n.detik % 3600
	n.menit = n.detik / 60
	n.detik = n.detik % 60

	fmt.Println(n.jam, n.menit, n.detik)
}
