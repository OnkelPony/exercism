package yacht

var categoryValues = map[string]int{
	"ones":   1,
	"twos":   2,
	"threes": 3,
	"fours":  4,
	"fives":  5,
	"sixes":  6,
}

func Score(dice []int, category string) int {
	switch category {
	case "yacht":
		return yachtValue(dice)
	case "choice":
		return choiceValue(dice)
	case "four of a kind":
		return fourOfAKindValue(dice)
	default:
		categoryValue, ok := categoryValues[category]
		if ok {
			return dicesValue(dice, categoryValue)
		}
	}
	return 0
}

func fourOfAKindValue(dice []int) int {
	result := make(map[int]int)
	for _, v := range dice {
		result[v]++
	}
	for v, c := range result {
		if c >= 4 {
			return 4 * v
		}
	}
	return 0
}

func choiceValue(dice []int) int {
	var result int
	for _, v := range dice {
		result += v
	}
	return result
}

func yachtValue(dice []int) int {
	for i := 1; i < len(dice); i++ {
		if dice[i] != dice[0] {
			return 0
		}
	}
	return 50
}

func dicesValue(dice []int, categoryValue int) int {
	var result int
	for _, v := range dice {
		if v == categoryValue {
			result += v
		}
	}
	return result
}
