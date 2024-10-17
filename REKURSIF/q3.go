package main

import "fmt"

func multip(x, y, total int) int {
	if y == 0 {
		return total
	} else {
		return multip(x, y-1, total+x)
	}
}

func main() {
	var x, y, total int
	fmt.Scan(&x, &y)
	fmt.Println(multip(x, y, total))
}
