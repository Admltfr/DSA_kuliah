package main

import "fmt"

type DataHarmonis struct {
	JumlahKebalikan float64
	JumlahBilangan  int
}

func rataHarmonis(n int) float64 {
	var data DataHarmonis
	var i int
	var input float64
	data.JumlahBilangan = n
	for i = 1; i <= n; i++ {
		fmt.Scan(&input)
		data.JumlahKebalikan = data.JumlahKebalikan + (1 / input)
	}
	return float64(data.JumlahBilangan) / data.JumlahKebalikan
}

func main() {
	var rata float64
	var n int
	fmt.Scan(&n)
	rata = rataHarmonis(n)
	fmt.Printf("%.2f\n", rata)
}
