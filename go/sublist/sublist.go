package sublist

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
	if isEqual(l1, l2) {
		return RelationEqual
	} else if isSublist(l1, l2) {
		return RelationSublist
	} else if isSuperlist(l1, l2) {
		return RelationSuperlist
	} else {
		return RelationUnequal
	}
}

func isEqual(l1, l2 []int) bool {
	if len(l1) != len(l2) {
		return false
	}
	for i := range l1 {
		if l1[i] != l2[i] {
			return false
		}
	}
	return true
}

func isSublist(l1, l2 []int) bool {
	equals := false
	short := len(l1)
	long := len(l2)
	if short > long {
		return false
	}	
	if short == 0 {
		return true
	}
	for o := 0; o <= long-short; o++ {
		for i := 0; i < short; i++ {
			if l1[i] == l2[o+i] {
				equals = true
			} else {
				equals = false
				break
			}
		}
		if equals {
			return equals
		}
	}
	return false
}

func isSuperlist(l1, l2 []int) bool {
	return isSublist(l2, l1)
}
