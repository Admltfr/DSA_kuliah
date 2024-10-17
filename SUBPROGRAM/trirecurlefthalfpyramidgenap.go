package main

import "fmt"

func printStars(N, i, j int, result string) string {
	var newResult string
	if i > N {
		return result
	}

	if i%2 == 0 {
		if j < i {
			newResult = result + "*"
			return printStars(N, i, j+1, newResult)
		} else {
			newResult = result + "\n"
			return printStars(N, i+1, 0, newResult)
		}
	}

	return printStars(N, i+1, 0, result)
}

func main() {
	var N int
	var result string
	fmt.Scan(&N)
	result = printStars(N, 1, 0, "")
	fmt.Print(result)
}
