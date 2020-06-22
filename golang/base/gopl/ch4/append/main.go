package main

import "fmt"

func main() {
	x := []int{1, 2, 3}
	y := []int{4, 5, 6}
	z := appendInt(x, y)
	fmt.Printf("%v\n", z)
}

func appendInt(x []int, y []int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 + len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}

	copy(z[len(x):], y)
	return z
}
