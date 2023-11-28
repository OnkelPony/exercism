package foodchain

import "fmt"

var firstVerse = "I know an old lady who swallowed a %s.\n"
var lastVerse = "I don't know why she swallowed the fly. Perhaps she'll die."
var uniSwallow = "She swallowed the %s to catch the %s%s.\n"
var onlyBirdSwallow = ""
var animals = []string{"fly", "spider", "bird", "cat", "dog", "goat", "cow", "horse"}
var swallows = []string{
	"It wriggled and jiggled and tickled inside her.\n",
	"How absurd to swallow a bird!\n",
	"Imagine that, to swallow a cat!\n",
	"What a hog, to swallow a dog!\n",
	"Just opened her throat and swallowed a goat!\n",
	"I don't know how she swallowed a cow!\n",
}

func Verse(v int) string {
	result := fmt.Sprintf(firstVerse, animals[v-1])
	if v == 8 {
		result += "She's dead, of course!"
		return result
	}
	if v > 1 {
		result += swallows[v-2]
		for i := v - 1; i > 0; i-- {
			if i == 2 {
				onlyBirdSwallow = " that wriggled and jiggled and tickled inside her"
			}
			result += fmt.Sprintf(uniSwallow, animals[i], animals[i-1], onlyBirdSwallow)
			onlyBirdSwallow = ""
		}
	}
	result += lastVerse
	return result
}

func Verses(start, end int) string {
	var result string
	for i := start; i <= end; i++ {
		result += Verse(i)
		if i != end {
			result += "\n\n"
		}
	}
	return result
}

func Song() string {
	return Verses(1, 8)
}
