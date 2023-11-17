package main

import (
	"fmt"
	"strings"
)

func main() {
	var S string
	fmt.Scan(&S)

	const d = "dream"
	const dr = "dreamer"
	const e = "erase"
	const er = "eraser"
	var size int
	for len(S) > 0 {
		size = 0
		if strings.HasSuffix(S, dr) {
			size = len(dr)
		} else if strings.HasSuffix(S, er) {
			size = len(er)
		} else if strings.HasSuffix(S, d) {
			size = len(d)
		} else if strings.HasSuffix(S, e) {
			size = len(e)
		} else {
			fmt.Println("NO")
			return
		}
		S = S[0 : len(S)-size]
	}
	fmt.Println("YES")
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
