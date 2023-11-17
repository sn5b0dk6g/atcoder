package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	ns := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &ns[i])
	}
	fmt.Print(f(ns))
}

func f(ns []int) int {
	for i, n := range ns {
		if n%2 != 0 {
			return 0
		}
		ns[i] = n / 2
	}
	return 1 + f(ns)
}
