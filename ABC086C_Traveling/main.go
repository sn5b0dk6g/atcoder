package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)
	pT, pX, pY := 0, 0, 0
	for i := 0; i < N; i++ {
		var t, x, y int
		fmt.Scanf("%d %d %d", &t, &x, &y)
		t, x, y = t-pT, x-pX, y-pY
		if x < 0 {
			x = x * -1
		}
		if y < 0 {
			y = y * -1
		}
		if t < x+y || t%2 != (x+y)%2 {
			fmt.Println("No")
			return
		}
		pT, pX, pY = t, x, y
	}
	fmt.Println("Yes")
}
