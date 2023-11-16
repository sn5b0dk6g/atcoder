package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)
	m := map[int][]int{}
	for i := 0; i < N; i++ {
		txy := make([]int, 3)
		for j := 0; j < 3; j++ {
			fmt.Scanf("%d", &txy[j])
		}
		m[i] = txy
	}
	var t, x, y int
	for i, txy := range m {
		if i == 0 {
			t = txy[0]
			x = txy[1]
			y = txy[2]
		} else {
			t = txy[0] - t
			x = txy[1] - x
			y = txy[2] - y
			if x < 0 {
				x = x * -1
			}
			if y < 0 {
				y = y * -1
			}
		}
		if t < x+y || (t%2 != 0 && x == y) || (t-1 == x+y) {
			//if t < x+y || t%2 != (x+y)%2 {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
