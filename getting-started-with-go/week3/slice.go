package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	s := make([]int, 0, 3)
	var inp string

	for {
		fmt.Print("Enter an integer: ")
		fmt.Scan(&inp)

		if inp == "x" {
			break
		}
		num, _ := strconv.Atoi(inp)

		s = append(s, num)
		sort.Ints(s)
		fmt.Println(s)
	}
}
