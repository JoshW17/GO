package main

import "fmt"

func sort(unsorted []int) ([]int, int) {
	finished := 0

	for i := 1; i < len(unsorted); i++ {
		if unsorted[i] < unsorted[i-1] {
			unsorted[i-1], unsorted[i] = unsorted[i], unsorted[i-1]
			finished++
		}
	}
	return unsorted, finished
}

func main() {
	sorted := []int{2, 3, 4, 3, 8, 7, 9, 5, 4, 6, 3, 6, 3}
	done := 1

	for done != 0 {
		sorted, done = sort(sorted)
	}

	fmt.Println(sorted)
}
