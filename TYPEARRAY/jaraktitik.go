package main

import "fmt"

type titik struct {
	x float64
	y float64
}

/*
Program jaraktitik
kamus
	type titik <
		x, y : real
	>
	p1, p2 : titik
Algoritma
	input(p1.x, p1.y, p2.x, p2.y)
	output(jarak(p1, p2))

Endprogram

function akar(x : real) real {
kamus
	y, z : real
algoritma
	y = x
	z = (y + (x / y)) / 2
	while y-z >= 0.00001  do
		y = z
		z = (y + (x / y)) / 2
	endwhile
	return z
endfunction
}

function jarak(p1, p2 : titik) real {
kamus
	jartik : real
algoritma
	jartik <- (p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y)
	return akar(jartik)
endfunction
}


*/

func jarak(p1, p2 titik) float64 {
	var dist float64
	dist = (p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y)
	return akar(dist)
}

func akar(x float64) float64 {
	var y, z float64
	y = x
	z = (y + (x / y)) / 2
	for y-z >= 0.00001 {
		y = z
		z = (y + (x / y)) / 2
	}
	return z
}

func main() {
	var p1, p2 titik
	fmt.Scan(&p1.x, &p1.y, &p2.x, &p2.y)
	fmt.Println(jarak(p1, p2))
}
