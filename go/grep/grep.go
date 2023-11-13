package grep

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Params struct {
	numberedLines bool
	listFilenames bool
	insensitiveCase bool
	viceVersa bool
	xCompleteLine bool
	moreFiles bool
}

func Search(pattern string, flags, files []string) []string {
	result := make([]string, 0)
	p  := Params{
		numberedLines: slices.Contains(flags, "-n"),
		listFilenames: slices.Contains(flags, "-l"),
		insensitiveCase: slices.Contains(flags, "-i"),
		viceVersa: slices.Contains(flags, "-v"),
		xCompleteLine: slices.Contains(flags, "-x"),
		moreFiles: len(files) > 1,
	}
	for _, f := range files {
		found := findInFile(pattern, p, f)
		result = append(result, found...)
	}
	return result
}

func findInFile(pattern string, p Params, f string) []string {
	var result []string
	var outLine string
	file, _ := os.Open(f)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var matches bool
	for lineNum := 1; scanner.Scan(); lineNum++ {
		inLine := scanner.Text()
		if p.insensitiveCase {
			inLine = strings.ToLower(inLine)
			pattern = strings.ToLower(pattern)
		}
		if p.xCompleteLine {
			matches = inLine == pattern
		} else {
			matches = strings.Contains(inLine, pattern)
		}
		// If the -v flag is set, it will take "unmatched" lines and if not, then "matched" ones
		if p.viceVersa != matches {
			if p.listFilenames {
				result = append(result, f)
				break
			} else if p.numberedLines {
				outLine = fmt.Sprintf("%d:%s", lineNum, scanner.Text())
			} else {
				outLine = scanner.Text()
			}
			if p.moreFiles {
				outLine = fmt.Sprintf("%s:%s", f, outLine)
			}
			result = append(result, outLine)
		}
	}
	return result
}
