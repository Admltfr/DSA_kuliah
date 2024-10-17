package main

import "fmt"

func main() {
	var bilangan int
	fmt.Scan(&bilangan)
	fmt.Println(isPrime(bilangan))

}

func isPrime(bil int) bool {
	/*  function mengembalikan boolean true jika prima atau false jika bukan bilangan prima */
	var count, i int
	var prime bool
	for i = 1; i <= bil; i++ {
		if i%2 == 0 {
			count++
		}
	}
	prime = count%2 == 0
	return prime
}
