package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	num1Str, err := rd.ReadString('\n')
	if err != nil {
		panic(err)
	}
	num1Str = strings.TrimSpace(num1Str)
	num2Str, err := rd.ReadString('\n')
	if err != nil {
		panic(err)
	}
	num2Str = strings.TrimSpace(num2Str)

	num1, err := strconv.Atoi(num1Str)
	if err != nil {
		panic(err)
	}
	num2, err := strconv.Atoi(num2Str)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s + %s = %d\n", num1Str, num2Str, num1+num2)

}
