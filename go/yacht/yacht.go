package yacht

import "reflect"

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
	case "full house":
		return fullHouseValue(dice)
	case "little straight":
		straight := map[int]int{1: 1, 2: 1, 3: 1, 4: 1, 5: 1}
		return straightValue(dice, straight)
	case "big straight":
		straight := map[int]int{2: 1, 3: 1, 4: 1, 5: 1, 6: 1}
		return straightValue(dice, straight)
	default:
		categoryValue, ok := categoryValues[category]
		if ok {
			return dicesValue(dice, categoryValue)
		}
	}
	return 0
}

func straightValue(dice []int, straight map[int]int) int {
	diceCombination := make(map[int]int)
	for _, v := range dice {
		diceCombination[v]++
	}
	if reflect.DeepEqual(diceCombination, straight) {
		return 30
	}
	return 0
}

func fullHouseValue(dice []int) int {
	var result int
	diceCombination := make(map[int]int)
	for _, v := range dice {
		diceCombination[v]++
	}
	if len(diceCombination) != 2 {
		return 0
	}
	for v, c := range diceCombination {
		if c == 3 {
			result += 3 * v
		}
		if c == 2 {
			result += 2 * v
		}
	}
	return result
}

func fourOfAKindValue(dice []int) int {
	diceCombination := make(map[int]int)
	for _, v := range dice {
		diceCombination[v]++
	}
	for v, c := range diceCombination {
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
