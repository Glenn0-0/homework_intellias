package main

import (
	"fmt"
	"sync"
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

	var wg sync.WaitGroup

	for i := 0; i < len(n); i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			fmt.Printf("slice %v: %v\n", i, sum(n[i]))
		}(i)
	}
	
	wg.Wait()
}