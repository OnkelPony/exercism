package main

import (
	"fmt"
)

func main() {
	fmt.Println("Print substrings")
	podretezce("abc", "", 0)
}

func podretezce(retezec, vysledek string, index int) {
	if index == len(retezec) {
		fmt.Println(vysledek)
	} else {
		podretezce(retezec, vysledek, index+1)

		// pravý podstrom - znak použijeme
		vysledek += string(retezec[index])
		podretezce(retezec, vysledek, index+1)
	}
}
