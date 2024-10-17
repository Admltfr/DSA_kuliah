/*package main

import "fmt"

const NMax int = 100

type tabChar [NMax]rune

func masukanArray(T *tabChar, n *int) {
	var char rune
	fmt.Scanf("%c", &char)
	for char != '.' {
		(*T)[*n] = char
		*n++
		fmt.Scanf("%c", &char)
	}
}

func lowerCase(T *tabChar, n int) {
	for i := 0; i < n; i++ {
		if 'A' <= (*T)[i] && (*T)[i] <= 'Z' {
			(*T)[i] = (*T)[i] + 'a' - 'A'
		}
	}
}

func cetakArray(T tabChar, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", T[i])
	}
}

func main() {
	var T tabChar
	var n int
	masukanArray(&T, &n)
	lowerCase(&T, n)
	cetakArray(T, n)
}
*/