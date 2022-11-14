// Package strand implements methods for making complementary RNA strand
package strand

// ToRNA returns RNA strand, complementary to given DNA strand
func ToRNA(dna string) string {
	dnaRnaMap := map[rune]rune{'G' : 'C', 'C' : 'G', 'T' : 'A', 'A' : 'U'}
	rna := ""
	for _, nucleotide := range dna {
		rna += string(dnaRnaMap[nucleotide])
	}
	return rna
}
