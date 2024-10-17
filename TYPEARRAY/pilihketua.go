package main

import "fmt"

type siswa struct {
	nama   string
	jumlah int
}

func main() {
	var s1, s2, s3 siswa
	var suara, i int
	fmt.Scan(&s1.nama, &s2.nama, &s3.nama)
	for i = 1; i <= 8; i++ {
		fmt.Scan(&suara)
		if suara == 1 {
			s1.jumlah++
		} else if suara == 2 {
			s2.jumlah++
		} else if suara == 3 {
			s3.jumlah++
		}
	}
	if s1.jumlah > s2.jumlah && s1.jumlah > s3.jumlah {
		fmt.Println(s1.nama, s1.jumlah)
	} else if s2.jumlah > s1.jumlah && s2.jumlah > s3.jumlah {
		fmt.Println(s2.nama, s2.jumlah)
	} else if s3.jumlah > s2.jumlah && s3.jumlah > s1.jumlah {
		fmt.Println(s3.nama, s3.jumlah)
	}
}
