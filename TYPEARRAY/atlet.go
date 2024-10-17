package main

import "fmt"

type atlet struct {
	negara string
	waktu  float64
}

func main() {
	var a1, a2, a3 atlet
	fmt.Scan(&a1.negara, &a1.waktu)
	fmt.Scan(&a2.negara, &a2.waktu)
	fmt.Scan(&a3.negara, &a3.waktu)
	pemenang(&a1, &a2, &a3)
}

func pemenang(a1, a2, a3 *atlet) {
	/* I.S.: Terdefinisi atlet1, atlet2, dan atlet3
	   F.S.: Mencetak string "Atlet asal ____ menang" jika salah satu dari 3 pelari tersebut menang atau "seri" jika catatan waktu ketiga pelari sama */
	if a1.waktu > a2.waktu && a1.waktu > a3.waktu {
		fmt.Printf("Atlet asal %s menang", a1.negara)
	} else if a2.waktu > a1.waktu && a2.waktu > a3.waktu {
		fmt.Printf("Atlet asal %s menang", a2.negara)
	} else if a3.waktu > a2.waktu && a3.waktu > a1.waktu {
		fmt.Printf("Atlet asal %s menang", a3.negara)
	} else if a1.waktu == a2.waktu && a2.waktu == a3.waktu {
		fmt.Print("seri")
	}
}
