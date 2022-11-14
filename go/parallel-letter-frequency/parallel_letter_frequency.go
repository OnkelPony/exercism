// Package letter implements methods to count frequency of runes in text,
// both sequentially and concurrently.
package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts frequency of each rune concurrently in string slices and returns this
// data as a FreqMap.
func ConcurrentFrequency(texts []string) FreqMap {
	result := make(FreqMap)
	freqs := make(chan FreqMap)
	for _, text := range texts {
		go func(text string) {
			freqs <- Frequency(text)
		}(text)
	}

	for range texts {
		freq := <-freqs
		for k, v := range freq {
			result[k] += v
		}
	}
	return result
}
