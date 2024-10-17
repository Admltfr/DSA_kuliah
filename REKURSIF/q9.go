package main

import "fmt"

func kali(n, m int) int {
	if m == 0 {
		return m
	} else {
		return n + kali(n, m-1)
	}
}

//function kali(n,m : integer) -> integer
//{mengembalikan bilangan bukat hasil penambahan n sebanyak m kali}
//	if n == 0 or m == 0 then
//		return m
//	else
//		return n + kali(n,m-1)
//	endif
//endfunction

//program perkalian
//kamus
//	n, m : integer
//algoritma
//	input(n,m)
//	output(kali(n,m))
//endprogram

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	fmt.Println(kali(n, m))

}
