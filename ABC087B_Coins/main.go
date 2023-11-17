package main

import (
	"fmt"
)

func main() {
	var A, B, C, X int
	fmt.Scanf("%d", &A)
	fmt.Scanf("%d", &B)
	fmt.Scanf("%d", &C)
	fmt.Scanf("%d", &X)
	fmt.Println(f(A, B, C, X))
}

func f(A, B, C, X int) int {
	t := 0

	for a := 0; a <= A; a++ {
		for b := 0; b <= B; b++ {
			for c := 0; c <= C; c++ {
				if 500*a+100*b+50*c == X {
					t++
				}
			}
		}
	}

	return t
}
