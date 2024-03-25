package main

import (
	"fmt"
	"os"
)

func main() {
	vypisObsahAdresare("..", "	")
}

func vypisObsahAdresare(cesta, odsazeni string) {
	soubor, err := os.Open(cesta)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer soubor.Close()

	stat, err := soubor.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	if stat.IsDir() {
		podadresare, err := os.ReadDir(cesta)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, podadresar := range podadresare {
			fmt.Println(odsazeni, podadresar.Name())
			if podadresar.IsDir() {
				vypisObsahAdresare(cesta+"/"+podadresar.Name(), odsazeni+"  ")
			}
		}
	} else {
		fmt.Println(odsazeni, soubor.Name())
	}
}
