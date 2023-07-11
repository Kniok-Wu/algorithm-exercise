package main

import "fmt"

func myPow(x float64, n int) float64 {
	flag := false

	var res float64 = 1

	if n < 0 {
		n *= -1
		flag = true
	}

	for n > 0 {
		if n%2 == 1 {
			res *= x
		}

		x *= x

		n >>= 1
	}

	if flag {
		res = 1 / res
	}
	return res
}

func main() {
	fmt.Println(myPow(2, 3))
}
