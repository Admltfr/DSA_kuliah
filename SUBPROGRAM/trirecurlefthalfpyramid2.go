package main

import "fmt"

func printstars2(N, i, j int, result string) string {
	var newResult string
	if i > N {
		return result
	}

	if i <= N {
		if j < i {
			newResult = result + "*"
			return printstars2(N, i, j+1, newResult)
		} else {
			newResult = result + "\n"
			return printstars2(N, i+1, 0, newResult)
		}
	}

	return printstars2(N, i+1, 0, result)
}

func main() {
	var N int
	var result string
	fmt.Scan(&N)
	result = printstars2(N, 1, 0, "")
	fmt.Print(result)
}
