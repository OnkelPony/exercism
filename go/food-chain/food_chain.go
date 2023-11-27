package foodchain

import "fmt"

var firstVerse = "I know an old lady who swallowed a %s.\n"
var lastVerse = "I don't know why she swallowed the fly. Perhaps she'll die."
var uniSwallow = "She swallowed the %s to catch the %s."
//var onlyBirdSwallow = "She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her."
var animals = []string{"fly", "spider", "bird", "cat", "dog", "goat", "cow", "horse"}
var swallows = []string{
	"It wriggled and jiggled and tickled inside her.",
	"How absurd to swallow a bird!",
	"Imagine that, to swallow a cat!",
	"What a hog, to swallow a dog!",
	"Just opened her throat and swallowed a goat!",
	"I don't know how she swallowed a cow!",
}

func Verse(v int) string {
	result := fmt.Sprintf(firstVerse, v-1)
	if v == 8 {
		result += "She's dead, of course!"
		return result
	}
	if v > 1 {
		result += swallows[v-2]
		for i := 0; i < v; i++ {
			result += fmt.Sprintf(uniSwallow, animals[i+1], animals[i])
		}
	}
	result += lastVerse
	return result
}

func Verses(start, end int) string {
	panic("Please implement the Verses function")
}

func Song() string {
	panic("Please implement the Song function")
}
