package main

import "fmt"

const NMAX int = 10

type tabInt [NMAX]int

func main() {
	var data tabInt
	var nData int
	baca(&data, &nData)
	cetak(data, nData)
	fmt.Println(jumlah(data, nData), rata_rata(data, nData))
}

func baca(A *tabInt, n *int) {
	/*IS: Array A dan n terdefinisi sembarang
	Proses: Setiap bilangan yang dibaca ditampung dalam sebuah variabel. Jika bilangan negatif, maka ubah menjadi bilangan positif dan masukan ke dalam array. Pembacaan berakhir jika terbaca bilangan 0.
	FS: Array A sebanyak n elemen berisi bilangan positif*/

	var i, input int
	i = 0
	fmt.Scan(&input)
	*n = 0
	for input != 0 {
		if *n < 10 {
			(*A)[i] = input
		}
		i++
		*n++
		fmt.Scan(&input)
	}
	if *n > NMAX {
		*n = NMAX
	}
}

func cetak(A tabInt, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Printf("%d ", A[i])
	}
	fmt.Println()
}

func jumlah(A tabInt, n int) int {
	/* Mengembalikan jumlah dari nilai bilangan pada elemen array A */
	var i, jumlah int

	for i = 0; i < n; i++ {
		if A[i] < 0 {
			jumlah = jumlah + A[i]*-1
		} else {
			jumlah = jumlah + A[i]
		}
	}
	return jumlah
}
func rata_rata(A tabInt, n int) float64 {
	/* Mengembalikan rata-rata bilangan */
	var rata float64
	rata = float64(jumlah(A, n)) / float64(n)
	return rata
}
