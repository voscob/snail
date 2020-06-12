package main

import (
	"fmt"

	"github.com/snail/services"
)

func main() {
	input := [][]int{
		{1, 2, 3, 1},
		{4, 5, 6, 4},
		{7, 8, 9, 7},
		{7, 8, 9, 7},
	}

	fmt.Println(services.Untwist(&input))
}
