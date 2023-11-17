package main

import (
	"fmt"
)

func main() {
	var N, A, B int
	fmt.Scanf("%d %d %d", &N, &A, &B)
	var t int
	for i := 1; i <= N; i++ {
		if sum := f(i); sum >= A && sum <= B {
			t += i
		}
	}
	fmt.Println(t)
}

func f(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}
