// Package protein implements methods to translate RNA to proteins
package protein

import "errors"

var codonProteins = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

// ErrStop is thrown, when stop codon appears
var ErrStop = errors.New("Protein creation stopped")

// ErrInvalidBase is thrown, when invalid base appears
var ErrInvalidBase = errors.New("Protein creation error - invalid base")

// FromCodon returns name of proteine, which codon is entered as argument
func FromCodon(input string) (string, error) {
	result := codonProteins[input]
	if result == "STOP" {
		return "", ErrStop
	}
	if result == "" {
		return result, ErrInvalidBase
	}
	return result, nil
}

// FromRNA returns list of proteines, which RNA is entered as argument
func FromRNA(input string) ([]string, error) {
	var result = []string{}
	codonLength := 3

	codon := ""
	for _, acid := range input {
		codon += string(acid)
		if len(codon) == codonLength {
			protein, err := FromCodon(codon)
			if err == ErrStop {
				return result, nil
			}
			if err == ErrInvalidBase {
				return result, err
			}
			result = append(result, protein)
			codon = ""
		}
	}

	return result, nil
}
