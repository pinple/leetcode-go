package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	n := 0
	for x != 0 {
		if n < math.MinInt32/10 || n > math.MaxInt32/10 {
			return 0
		}
		n = n*10 + x%10
		x = x / 10
	}
	return n
}

func main() {
	fmt.Println(reverse(123))
}
