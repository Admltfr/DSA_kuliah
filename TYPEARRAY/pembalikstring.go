package main

import "fmt"

const N int = 100

type tab [N]rune

func masukanArray(T *tab, n *int) {
	var char rune
	fmt.Scanf("%c", &char)
	for char != '.' {
		(*T)[*n] = char
		*n++
		fmt.Scanf("%c", &char)
	}
}

func reverseArray(T *tab, n int) {
	var temp rune
	var i, j int
	j = n - 1
	for i = 0; i < j; i++ {
		temp = (*T)[i]
		(*T)[i] = (*T)[j]
		(*T)[j] = temp
		j--
	}
}

func cetakArray(T tab, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", T[i])
	}
}

func main() {
	var T tab
	var n int
	masukanArray(&T, &n)
	reverseArray(&T, n)
	cetakArray(T, n)
}
