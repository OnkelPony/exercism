package lasagna

const (
	noodlesLayer = 50
	sauceLayer   = 0.2
)

// PreparationTime returns the time to cook lasagna based on the number of layers.
func PreparationTime(layers []string, t int) int {
	if t == 0 {
		t = 2
	}
	return t * len(layers)
}

// Quantities returns the amounts of noodles and sauce, based on the number of layers.
func Quantities(layers []string) (int, float64) {
	var nCount, sCount int
	for _, layer := range layers {
		switch layer {
		case "noodles":
			nCount++
		case "sauce":
			sCount++
		}
	}
	return noodlesLayer * nCount, sauceLayer * float64(sCount)
}

// AddSecretIngredient takes two lists as parameter and returns second list with the last item from the first list at the end
func AddSecretIngredient(friendList, myList []string) []string {
	return append(myList, friendList[len(friendList)-1])
}

// ScaleRecipe returns slice of float64 of the amounts for the desired number of portions.
func ScaleRecipe(amounts []float64, portions int) []float64 {
	result := make([]float64, len(amounts))
	for i, amount := range amounts {
		result[i] = float64(portions) / 2 * amount
	}
	return result
}
