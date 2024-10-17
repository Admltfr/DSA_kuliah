package main

import "fmt"

type geometry struct {
	area      int
	perimeter int
}

type rectangle struct {
	length   int
	width    int
	color    string
	property geometry
}

/*
program persegipanjang
kamus
	type geometry <
		area, perimeter : integer
	>

	type rectangle <
		length, width : integer
		color : string
		property : geometry
	>

	data : rectangle

algoritma
	isidata(data)
	hitung(data)
	output(data.property.area, data.property.perimeter)
endprogram

procedure isidata(in/out persegi : rectangle)
algoritma
	input(persegi.length, persegi.width, persegi.color)
endprocedure

procedure hitung(in/out persegi : rectangle)
algoritma
	persegi.property.area = persegi.length * persegi.width
	persegi.property.perimeter = (persegi.length * 2) + (persegi.width * 2)
endprocedure
*/

func isidata(persegi *rectangle) {
	fmt.Scan(&persegi.length, &persegi.width, &persegi.color)
}

func hitung(persegi *rectangle) {
	persegi.property.area = persegi.length * persegi.width
	persegi.property.perimeter = (persegi.length * 2) + (persegi.width * 2)
}

func main() {
	var data rectangle
	isidata(&data)
	hitung(&data)
	fmt.Println(data.property.area, data.property.perimeter)
}
