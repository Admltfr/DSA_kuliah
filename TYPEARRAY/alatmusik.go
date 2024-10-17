package main

import "fmt"

type ins struct {
	jenis, merek, tipe string
	harga              int
}

func main() {
	var alat1, alat2, alat3 ins
	var inputharga int
	var inputjenis, inputmerek, inputtipe string
	fmt.Scan(&alat1.jenis, &alat1.merek, &alat1.tipe, &alat1.harga)
	fmt.Scan(&alat2.jenis, &alat2.merek, &alat2.tipe, &alat2.harga)
	fmt.Scan(&alat3.jenis, &alat3.merek, &alat3.tipe, &alat3.harga)
	for inputjenis != "selesai" {
		fmt.Scan(&inputjenis)
		if inputjenis == "selesai" {
			break
		}
		fmt.Scan(&inputmerek, &inputtipe, &inputharga)
		if inputjenis == alat1.jenis {
			if inputharga < alat1.harga {
				alat1.merek = inputmerek
				alat1.tipe = inputtipe
				alat1.harga = inputharga

			}
		} else if inputjenis == alat2.jenis {
			if inputharga < alat2.harga {
				alat2.merek = inputmerek
				alat2.tipe = inputtipe
				alat2.harga = inputharga

			}
		} else if inputjenis == alat3.jenis {
			if inputharga < alat3.harga {
				alat3.merek = inputmerek
				alat3.tipe = inputtipe
				alat3.harga = inputharga

			}
		}
	}
	fmt.Println(alat1.jenis, alat1.merek, alat1.tipe, alat1.harga)
	fmt.Println(alat2.jenis, alat2.merek, alat2.tipe, alat2.harga)
	fmt.Println(alat3.jenis, alat3.merek, alat3.tipe, alat3.harga)
}
