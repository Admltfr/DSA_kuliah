package main

import "fmt"

const NMAX int = 1000

type tabreal [n]float64

func main() {
	var arr tabreal
	inputarray(&arr)

}

func inputarray(a *tabreal) {
	var i int
	i = 0
	for i < n {
		fmt.Scan(&a[i])
		i++
	}
}

func findmax(a tabreal) float64 {
	var i int
	var max float64
	i = 0
	max = a[0]
	for i < n {
		if a[i] > max {
			max = a[i]
		}
		i++
	}
	return max
}

func findmin(a tabreal) float64 {
	var i int
	var min float64
	i = 0
	min = a[0]
	for i < n {
		if a[i] < min {
			min = a[i]
		}
		i++
	}
	return min
}

/*
Program minmax
kamus
	constant n : integer = 6
	type tabreal : array [1..n] of real
	arr : tabreal
algoritma
	inputarray(arr)
	output("Nilai Maksimum:", findmax(arr))
	output("Nilai Minimum:", findmin(arr))
endprogram

Procedure inputarray(in/out a : tabreal)
{I.S. -
F.S. array a telah terisi oleh nilai yang dimasukkan}
kamus
	i : integer
algoritma
	i <- 1
	while i <= n do
		input(a[i])
		i <- i + 1
	endwhile
endprocedure

Function findmax(a tabreal) -> real
{mengembalikan nilai maksimum dari array a}
kamus
	i : integer
	max : real
algoritma
	i <- 1
	max = a[1]
	while i <= n do
		if a[i] > max then
			max = a[i]
		endif
		i <- i + 1
	endwhile
	return max
endfunction

Function findmin(a tabreal) -> real
{mengembalikan nilai minimum dari array a}
kamus
	i : integer
	min : real
algoritma
	i <- 1
	min = a[1]
	while i <= n do
		if a[i] < min then
			min = a[i]
		endif
		i <- i + 1
	endwhile
	return min
endfunction
*/
