package main

import "fmt"

type sushi struct {
	t int
	d int64
}

func main() {
	var N, K int
	fmt.Scanf("%d %d", &N, &K)
	var tbp int
	var dbp int64
	choice := make([]sushi, K)
	m := map[int]int{}
	for i := 0; i < N; i++ {
		var t int
		var d int64
		fmt.Scanf("%d %d", &t, &d)
		if i > (K - 1) {
			choice[i].t = t
			choice[i].d = d
			m[t] += 1
			tbp = len(m) ^ 2
			dbp += d
		}
	}
}
