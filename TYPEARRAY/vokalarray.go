package main

import "fmt"

const N = 100

var n int
var s string
var arr [N]byte

func main() {
	fmt.Scan(&s)
	for n < len(s) {
		arr[n] = s[n]
		n++
	}
	fmt.Println(searchVokal(s))
}

func searchVokal(x string) bool {
	var count int
	var vokal bool
	n = 0
	for n < len(s) {
		if arr[n] == 'a' || arr[n] == 'i' || arr[n] == 'u' || arr[n] == 'e' || arr[n] == 'o' {
			count = count + 1
		}
		n++
	}
	if count > 0 {
		vokal = true
	}
	return vokal

}
