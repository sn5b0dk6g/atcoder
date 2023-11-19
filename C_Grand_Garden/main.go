package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)

	hs := make([]int, N)
	for i := range hs {
		fmt.Scanf("%d", &hs[i])
	}

	count := 0
	for i := 0; i < N; i++ {
		if i == 0 {
			count += hs[i]
		} else if hs[i] > hs[i-1] {
			count += hs[i] - hs[i-1]
		}
	}

	fmt.Printf("%d\n", count)
}
