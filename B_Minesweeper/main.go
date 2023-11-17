package main

import (
	"fmt"
	"strconv"
)

func main() {
	var H, W int
	fmt.Scanf("%d %d", &H, &W)
	matrix := make([][]int, H)
	for i := 0; i < H; i++ {
		matrix[i] = make([]int, W)
		var s string
		fmt.Scan(&s)
		runes := []rune(s)
		for j := 0; j < len(runes); j++ {
			var v int
			if string(runes[j]) == "#" {
				v = 9
			} else {
				v = 0
			}
			matrix[i][j] = v
		}
	}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if matrix[i][j] < 9 {
				continue
			}
			if i > 0 {
				matrix[i-1][j] += 1
				if j < W-1 {
					matrix[i-1][j+1] += 1
				}
				if j > 0 {
					matrix[i-1][j-1] += 1
				}
			}
			if i < H-1 {
				matrix[i+1][j] += 1
				if j > 0 {
					matrix[i+1][j-1] += 1
				}
				if j < W-1 {
					matrix[i+1][j+1] += 1
				}
			}
			if j < W-1 {
				matrix[i][j+1] += 1
			}
			if j > 0 {
				matrix[i][j-1] += 1
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		var p string
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] >= 9 {
				p += "#"
			} else {
				p += strconv.Itoa(matrix[i][j])
			}
		}
		fmt.Println(p)
	}
}
