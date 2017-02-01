package main

import (
	"fmt"
	"os"
	"strconv"
)

func slicesum(pascal []int) int {

	pascal = pascal[:len(pascal)-1]

	for len(pascal) > 1 {

		pascal[1] = pascal[0] + pascal[1]

		pascal = pascal[1:]

	}

	return pascal[0]

}

func main() {

	length := os.Args[1]

	pascal := make([]int, 0)

	reducingCount, _ := strconv.Atoi(length)

	for i := 0; i < reducingCount; i++ {

		pascal = append(pascal, 1)

	}

	fmt.Printf("%d ", 1)

	for len(pascal) > 1 {

		fmt.Printf("%d ", slicesum(pascal))
		pascal = pascal[:len(pascal)-1]

	}

	fmt.Println()

}
