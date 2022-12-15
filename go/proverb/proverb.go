// Package proverb implements method to assemble proverb from given words.
package proverb

// Proverb returns whole proverb, using given words to assemble it.
func Proverb(rhyme []string) []string {
	var result []string
	rhymeLength := len(rhyme)
	if rhymeLength == 0 {
		return result
	}
	if rhymeLength > 1 {
		for i, firstWord := range rhyme {
			secondWord := rhyme[i+1]
			result = append(result, "For want of a "+firstWord+" the "+secondWord+" was lost.")
			if i == rhymeLength-2 {
				break
			}
		}
	}
	result = append(result, "And all for the want of a"+" "+rhyme[0]+".")
	return result
}
