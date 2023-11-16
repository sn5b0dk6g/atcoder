package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)
	list := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &list[i])
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i] > list[j]
	})

	var alice, bob int
	for i, n := range list {
		if i%2 == 0 {
			alice += n
		} else {
			bob += n
		}
	}
	fmt.Println(alice - bob)
}
