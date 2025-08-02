package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var rows, cache [][]int

func main() {
	var x, y, result int
	fmt.Scan(&x, &y)
	scanner := bufio.NewScanner(os.Stdin)
	rows = make([][]int, x)

	for i := 0; i < x; i++ {
		scanner.Scan()
		line := scanner.Text()
		numbers := strings.Fields(line)
		rows[i] = make([]int, y)
		for j := 0; j < y; j++ {
			rows[i][j], _ = strconv.Atoi(numbers[j])
		}
	}
	cache = make([][]int, x)
	for i := 0; i < x; i++ {
		cache[i] = make([]int, y)
		for j := 0; j < y; j++ {
			cache[i][j] = -1
		}
	}
	for i := 0; i < y; i++ {
		sum := memorizedCountGame(i, 0, x, y)
		result = maxSum(sum, result)
	}
	fmt.Println(result)
}

func memorizedCountGame(row, line int, maxLine, maxRow int) int {
	if cache[line][row] != -1 {
		return cache[line][row]
	}

	result := rows[line][row]
	if line == maxLine-1 {
		cache[line][row] = result
		return result
	}

	var maxNext int
	if row > 0 {
		maxNext = memorizedCountGame(row-1, line+1, maxLine, maxRow)
	}
	maxNext = max(maxNext, memorizedCountGame(row, line+1, maxLine, maxRow))
	if row < maxRow-1 {
		maxNext = max(maxNext, memorizedCountGame(row+1, line+1, maxLine, maxRow))
	}

	result += maxNext
	cache[line][row] = result
	return result
}

func max(nums ...int) int {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func maxSum(sum, result int) int {
	if sum > result {
		return sum
	}
	return result
}

// 以下は再帰関数を使った解放
// 時間計算量がO(3^x*y)のためとても遅い

// func countGame(row, line, subtotal, maxLine, maxRow int) int {
// 	var left, center, right int
// 	subtotal += rows[line][row]

// 	if line == maxLine-1 {
// 		return subtotal
// 	}
// 	if row > 0 {
// 		left = countGame(row-1, line+1, subtotal, maxLine, maxRow)
// 	}
// 	center = countGame(row, line+1, subtotal, maxLine, maxRow)
// 	if row < maxRow-1 {
// 		right = countGame(row+1, line+1, subtotal, maxLine, maxRow)
// 	}
// 	return max(left, center, right)
// }
