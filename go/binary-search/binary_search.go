package binarysearch

// SearchInts returns index of the key in the list, in case of error it returns -1
func SearchInts(list []int, key int) int {
	h := len(list)
	l := 0
	for toTest := h; toTest > 0; toTest /= 2 {
		m := (l + h) / 2
		switch {
		case key > list[m]:
			l = m
		case key < list[m]:
			h = m
		default:
			return m
		}
	}
	return -1
}
