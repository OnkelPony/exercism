package series

func All(n int, s string) []string {
	if len(s) < n {
		return nil
	}
	var result []string
	howMany := len(s) - n + 1
	for i := 0; i < howMany; i++ {
		result = append(result, s[i:i+n])
	}
	return result
}

func UnsafeFirst(n int, s string) string {
	// var result string
	if len(s) < n {
		return ""
	}
	return s[:n]
}

func First(n int, s string) (first string, ok bool) {
	if len(s) < n {
		return "", false
	}
	return s[:n], true
}
