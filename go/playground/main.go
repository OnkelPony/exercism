package main

import (
	"fmt"
)

func main() {
	fmt.Println("Print bins")
	binCisla([]int{0, 0, 0}, 0)
}

func binCisla(pole []int, index int) {
	if index == len(pole) {
		fmt.Println(pole)
	} else {
		pole[index] = 0
		binCisla(pole, index+1)
		pole[index] = 1
		binCisla(pole, index+1)
	}
}
