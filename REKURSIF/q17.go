package main

import "fmt"

func strip(n int) {
	if n < 10 {
		fmt.Print(n)
	} else {
		fmt.Print(n%10, "-")

		strip(n / 10)
	}
}

//procedure strip(in n : integer)
//algoritma
//	if n < 10 then
//		output(n)
//	else
//		output(n mod 10,"-")
//		strip(n div 10)
//	endif
//endprocedure

//program cetakstrip
//kamus
//	n : integer
//algoritma
//	input(n)
//	strip(n)
//endprogram

func main() {
	var n int
	fmt.Scan(&n)
	strip(n)
}
