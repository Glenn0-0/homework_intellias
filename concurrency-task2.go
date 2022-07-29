package main

import (
	"fmt"
)

func sum(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func main() {
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}
	
	channel := make(chan int)
	var res int

	for _, row := range n {
		go calcSum(row, channel)
	}

	for i := 0; i < len(n); i++ {
		res += <- channel
	}

	fmt.Println(res)
}

func calcSum(arr []int, ch chan int) {
	var sum int
	for _, num := range arr {
		sum += num
	}

	ch <- sum
}