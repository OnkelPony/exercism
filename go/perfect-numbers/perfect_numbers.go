package perfect

import (
	"errors"
	"math"
)

type Classification int

const ClassificationAbundant Classification = 1

const ClassificationDeficient Classification = -1
const ClassificationPerfect Classification = 0

var ErrOnlyPositive = errors.New("Input can be only positive integer")

func Classify(n int64) (Classification, error) {
	if n == 1 {
		return ClassificationDeficient, nil
	}

	if n < 1 {
		return 0, ErrOnlyPositive
	}

	var total int64 = 1
	var i int64 = 2
	for ; i <= int64(math.Sqrt(float64(n))); i++ {
		j := float64(n) / float64(i)
		if j == math.Floor(j) && i != int64(j){
			total += i
			total += int64(j)
		}
	}
	if total == n {
		return ClassificationPerfect, nil
	} else if total > n {
		return ClassificationAbundant, nil
	}
	return ClassificationDeficient, nil
}
