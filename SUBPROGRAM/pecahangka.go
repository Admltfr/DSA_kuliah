package main

import "fmt"

func cacah(n int) {
	/* I.S. terdefinisi sebuah bilangan bulat "bilangan"
	   F.S. program mengeluarkan nilai yang menyatakan nilai setiap digit dari X dengan setiap nilai dipisahkan oleh spasi */
	if n > 0 {
		fmt.Print(n%10, " ")
		cacah(n / 10)
	}
}

func main() {
	var bilangan int
	fmt.Scan(&bilangan)
	cacah(bilangan)
}
