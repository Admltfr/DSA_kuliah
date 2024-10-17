package main

import "fmt"

func main() {
	type mahasiswa struct {
		nama string
		nim  string
		ipk  float64
	}
	var m1, m2, m3 mahasiswa
	var rata float64
	fmt.Scan(&m1.nama, &m1.nim, &m1.ipk)
	fmt.Scan(&m2.nama, &m2.nim, &m2.ipk)
	fmt.Scan(&m3.nama, &m3.nim, &m3.ipk)
	rata = (m1.ipk + m2.ipk + m3.ipk) / 3

	if m1.ipk > m2.ipk && m1.ipk > m3.ipk {
		fmt.Printf("%.2f\n%s", rata, m1.nama)
	} else if m2.ipk > m1.ipk && m2.ipk > m3.ipk {
		fmt.Printf("%.2f\n%s", rata, m2.nama)
	} else if m3.ipk > m2.ipk && m3.ipk > m1.ipk {
		fmt.Printf("%.2f\n%s", rata, m3.nama)
	}
}
