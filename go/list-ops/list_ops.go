package listops

type IntList []int
type binFunc func(x, y int) int
type predFunc func(n int) bool
type unaryFunc func(x int) int


func (l IntList) Length() int {
	length := 0
	for range l {
		length++
	}
	return length
}

func (l IntList) Foldr(fn binFunc, initial int) int {
	llength := l.Length() - 1
	for index := llength; index >= 0; index-- {
		initial = fn(l[index], initial)
	}
	return initial
}

func (l IntList) Foldl(fn binFunc, initial int) int {
	for _, number := range l {
		initial = fn(initial, number)
	}
	return initial
}

func (l IntList) Filter(fn predFunc) IntList {
	listLength := l.Length()
	if listLength == 0 {
		return l
	}
	index := 0
	for _, number := range l {
		if fn(number) {
			index++
		}
	}
		result := make (IntList, index)

		index = 0
	for _, number := range l {
		if fn(number) {
			result[index] = number
			index++
		}
	}
	return result
}

func (l IntList) Map(fn unaryFunc) IntList {
	listLength := l.Length()
	if listLength == 0 {
		return l
	}
	result := make(IntList, listLength)
	for index, number := range l {
		result[index] = fn(number)
	}
	return result
}

func (l IntList) Reverse() IntList {
	listLength := l.Length()
	if listLength == 0 {
		return l
	}
	index := 0

	result := make (IntList, listLength)
	for i := listLength - 1; i >= 0; i-- {
		result[index] = l[i]
		index++
	}
	return result
}

func (l IntList) Append(this IntList) IntList {
	listLength := l.Length()
	appendListLength := this.Length()
	if listLength == 0 {
		return this
	}
	if appendListLength == 0 {
		return l
	}
	result := make(IntList, listLength + appendListLength)
	for i, number := range l {
		result[i] = number
	}
	for i, number := range this {
		result[listLength + i] = number
	}
	return result
}

func (l IntList) Concat(args []IntList) IntList {
	listLength := l.Length()
	argsLength := len(args)
	if listLength == 0 && argsLength == 0 {
		return l
	}
	for _, listOfInts := range args {
		l = l.Append(listOfInts)
	}
	return l
}