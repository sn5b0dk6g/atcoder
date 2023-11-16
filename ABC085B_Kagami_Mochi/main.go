package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)
	m := make(map[int]bool)
	var key int
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &key)
		m[key] = true
	}
	fmt.Println(len(m))
}
