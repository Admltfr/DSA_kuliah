package main

import "fmt"

func kakiHewan(n int) int {
	if n == 1 {
		return 2
	} else if n%2 == 0 {
		return 3 + kakiHewan(n-1)
	} else {
		return 2 + kakiHewan(n-1)
	}
}

func main() {
	var N, result int
	fmt.Scan(&N)
	result = kakiHewan(N)
	fmt.Println(result)
}
