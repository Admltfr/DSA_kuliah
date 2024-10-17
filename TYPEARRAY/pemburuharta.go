package main

import "fmt"

type koordinat struct {
	x, y int
}

func main() {
	var k koordinat
	var i, kx, ky int
	fmt.Scan(&k.x, &k.y)
	i = 1
	for i <= 3 && (kx != k.x || ky != k.y) {
		fmt.Scan(&kx, &ky)
		if kx != k.x || ky != k.y {
			fmt.Println("Coba Lagi")
		}
	}
	if kx == k.x && ky == k.y {
		fmt.Println("Yay Menang")
	}
}
