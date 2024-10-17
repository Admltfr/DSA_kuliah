package main

import "fmt"

func main() {
	var nama, ucapan string
	fmt.Scan(&nama)
	ucapan = greeting(nama)
	fmt.Println(ucapan)

}

func greeting(s string) string {
	/* fungsi string membaca nama dan mengembalikan string "Halo, <nama>"*/
	return "Halo, " + s
}
