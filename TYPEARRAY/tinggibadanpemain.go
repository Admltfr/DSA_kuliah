package main

import "fmt"

const NMAX int = 11

type pemain struct {
	nama, nomorPunggung, posisi string
	tinggi                      int
}

// tabPemain adalah alias array pemain dengan maks elemen NMAX
type tabPemain [NMAX]pemain

func main() {
	var timnas tabPemain
	var nPemain int
	nPemain = 0
	// baca data
	bacaData(&timnas, &nPemain)
	// cetak data pemain
	cetakPemain(timnas, nPemain)
	// cetak nama pemain tertinggi
	cetakNamaPemainTertinggi(timnas, nPemain)
	// cetak nama pemain terendah
	cetakNamaPemainTerendah(timnas, nPemain)
}
func bacaData(A *tabPemain, n *int) {
	/*
		IS: Array A dengan banyak elemen n terdefinisi sembarang
		Proses: Membaca nama, nomor punggung, posisi, dan tinggi badan Jika array belum penuh dan nama bukan "none", maka nama, nomor punggung, pos dan tinggi badan dimasukkan ke dalam array.
		FS: Array A sebanyak n elemen berisi nilai
	*/
	var i int
	var nama string

	for i = 0; i < NMAX; i++ {
		fmt.Scan(&nama)
		if nama == "none" {
			break
		}
		A[i].nama = nama
		fmt.Scan(&A[i].nomorPunggung, &A[i].posisi, &A[i].tinggi)
		*n = *n + 1
	}
}
func cetakPemain(A tabPemain, n int) {
	/*
			IS: Array A dengan banyak elemen n terdefinisi
			FS: Tercetak di layar elemen array A sebanyak n dengan format:
			"Data pemain: <nama1> <nomorPunggung1> <posisi1> <tinggi1>
		<nama2> <nomorPunggung2> <posisi2> <tinggi2>
		<naman> <nomorPunggungn> <posisin> <tinggin>"
		...
	*/
	var i int
	fmt.Println("Data Pemain:")
	for i = 0; i < n; i++ {
		fmt.Println(A[i].nama, A[i].nomorPunggung, A[i].posisi, A[i].tinggi)
	}

}
func cetakNamaPemainTertinggi(A tabPemain, n int) {
	/* IS: Array A sebanyak n elemen terdefinisi
	   FS: Mencetak nama pemain dengan badan tertinggi dengan format: "Pemain dengan badan tertingggi: <nama>"
	   Asumsi: Hanya ada 1 pemain dengan badan tertinggi
	*/
	var i, max int

	max = 0
	for i = 0; i < n; i++ {
		if A[max].tinggi < A[i].tinggi {
			max = i
		}

	}
	fmt.Println("Pemain dengan badan tertinggi:", A[max].nama)

}
func cetakNamaPemainTerendah(A tabPemain, n int) {
	/* IS: Array A sebanyak n elemen terdefinisi
	   FS: Mencetak nama pemain dengan badan terendah dengan format: "Pemain dengan badan terendah: <nama>"" Asumsi: Hanya ada 1 pemain dengan badan terendah
	*/
	var i, min int

	min = 0
	for i = 0; i < n; i++ {
		if A[min].tinggi > A[i].tinggi {
			min = i
		}

	}
	fmt.Println("Pemain dengan badan terendah:", A[min].nama)

}
