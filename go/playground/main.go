package main

import (
	"fmt"
)

func main() {
	fmt.Println("Print signs")
	znamenka([]int{1, 2, 3}, 0)
}

func znamenka(cisla []int, index int) {
	if index == len(cisla) {
		sum := 0
		for i := 0; i < len(cisla); i++ {
			sum += cisla[i]
		}
		if sum == 0 {
			for i := 0; i < len(cisla); i++ {
				fmt.Printf("%d ", cisla[i])
			}
			fmt.Println()
		}
	} else {
		znamenka(cisla, index+1)
		cisla[index] *= -1
		znamenka(cisla, index+1)
	}
}
