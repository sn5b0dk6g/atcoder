package main

import (
	"fmt"
)

func main() {
	var N, Y int
	fmt.Scanf("%d %d", &N, &Y)

	for yukichi := 0; yukichi <= N; yukichi++ {
		for higuchi := 0; higuchi <= N-yukichi; higuchi++ {
			natsume := N - yukichi - higuchi
			total := yukichi*10000 + higuchi*5000 + natsume*1000
			if total == Y {
				fmt.Printf("%d %d %d\n", yukichi, higuchi, natsume)
				return
			}
		}
	}
	fmt.Println("-1 -1 -1")
}
