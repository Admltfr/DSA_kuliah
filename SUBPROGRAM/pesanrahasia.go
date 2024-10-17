package main

import "fmt"

func main() {
	var pesan string
	fmt.Scan(&pesan)
	fmt.Println(kirimPesanRahasia(pesan))
}
func kirimPesanRahasia(s string) string {
	// mengembalikan pesan rahasia pada karakter ke-1, ke-5, ke-9, dan ke-13 dari string masukan.
	if len(s) < 12 {
		s = ""
	}
	var hasil string
	hasil = s[0:1] + s[4:5] + s[8:9] + s[12:13]

	return hasil
}
