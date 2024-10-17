package main

import "fmt"

func findyear(s string, hasil *string) {
	*hasil = "20" + string(s[4]) + string(s[5])
}

func main() {
	var s, hasil string
	fmt.Scan(&s)
	findyear(s, &hasil)
	fmt.Println(hasil)
}
