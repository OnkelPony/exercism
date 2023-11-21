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
	default:
		categoryValue, ok := categoryValues[category]
		if ok {
			return dicesValue(dice, categoryValue)
		}
	}
	return 0
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
