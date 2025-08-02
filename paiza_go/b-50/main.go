package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var x int
	var y int
	fmt.Scan(&x)
	fmt.Scan(&y)

	scanner := bufio.NewScanner(os.Stdin)
	rows := make([]string, x)
	for i := 0; i < x; i++ {
		scanner.Scan()
		rows[i] = scanner.Text()
	}
	// fmt.Println(rows)
	bombed := make([]string, x)
	tate := make(map[int]bool, y)
	for i := 0; i < x; i++ {
		bombed[i] = rows[i]
		if strings.Contains(rows[i], "#") {
			bombed[i] = bomb(y)
			for j := 0; j < y; j++ {
				if rows[i][j] == '#' {
					if _, ok := tate[j]; !ok {
						tate[j] = true
					}
				}
			}
			continue
		}
	}
	for k := range tate {
		for l := 0; l < x; l++ {
			temp := []rune(bombed[l])
			temp[k] = '#'
			bombed[l] = string(temp)
		}
	}
	var bombedCount int
	for i := 0; i < x; i++ {
		if strings.Contains(bombed[i], "#") {
			bombedCount += strings.Count(bombed[i], "#")
		}
	}
	fmt.Println(bombedCount)
}

func bomb(y int) string {
	return strings.Repeat("#", y)
}
