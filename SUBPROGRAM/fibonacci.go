package main

import "fmt"

func fibo(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibo(n-1) + fibo(n-2)
	}
}

func cetakfibo(n int) {
	if n >= 0 {
		cetakfibo(n - 1)
		fmt.Print(fibo(n), " ")
	}

}

func main() {
	var n int
	fmt.Scan(&n)
	cetakfibo(n)
}
