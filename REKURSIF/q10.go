package main

import "fmt"

func jumlahdigit(n int) int {
	if n == 0 {
		return n
	} else {
		return n%10 + jumlahdigit(n/10)
	}
}

//function kali(n : integer) -> integer
//{mengembalikan bilangan bulat hasil penjumlahan digit n}
//		if n == 0 then
//				return n
//		else
//				return n mod 10 + jumlahdigit(n div 10)
//		endif
//endfunction

//program perkalian
//kamus
//		n : integer
//algoritma
//		input(n)
//		output(jumlahdigit(n))
//endprogram

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(jumlahdigit(n))

}
