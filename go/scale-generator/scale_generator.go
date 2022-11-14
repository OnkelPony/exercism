/*
Package scale implements generation of scales based on tonic and interval
*/
package scale

import "strings"

var sharpChrom = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
var flatChrom = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
var flatScales = []string{"F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb"}

//Returns scale based on tonic and interval
func Scale(tonic string, interval string) []string {
	generatedScale := make([]string, 0)
	useChrom := sharpChrom
	var positionOfTone int32
	isFlat, _ := find(flatScales, tonic)
	if isFlat {
		useChrom = flatChrom
	}
	tonic = strings.Title(tonic)

	// First note has no interval
	generatedScale = append(generatedScale, tonic)
	_, positionOfTone = find(useChrom, tonic)

	if interval != "" {
		generatedScale = makeScale(interval, positionOfTone, generatedScale, useChrom)
	} else {
		for range useChrom[:len(useChrom)-1] {
			positionOfTone += 1
			generatedScale = append(generatedScale, getTone(useChrom, positionOfTone))
		}
	}

	return generatedScale
}

func makeScale(interval string, positionOfTone int32, generatedScale []string, useChrom []string) []string {
	for _, step := range interval[:len(interval)-1] {
		if step == 'm' {
			step = 1
		} else if step == 'M' {
			step = 2
		} else if step == 'A' {
			step = 3
		}
		positionOfTone += step
		generatedScale = append(generatedScale, getTone(useChrom, positionOfTone))
	}
	return generatedScale
}

func getTone(chrom []string, positionOfTone int32) string {
	indexesOfChrom := int32(len(chrom))
	positionOfTone = positionOfTone % indexesOfChrom
	return chrom[positionOfTone]
}

func find(scales []string, tonic string) (bool, int32) {
	for i, tone := range scales {
		if tone == tonic {
			return true, int32(i)
		}
	}
	return false, -1
}
