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
	return false
}

func isSublist(l1, l2 []int) bool {
	return false
}

func isSuperlist(l1, l2 []int) bool {
	return false
}
