package grep

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func Search(pattern string, flags, files []string) []string {
	result := make([]string, 0)
	for _, f := range files {
		found := findInfile(pattern, flags, f)
		result = append(result, found...)
	}
	return result
}

func findInfile(pattern string, flags []string, f string) []string {
	var result []string
	var lineNum int
	file, _ := os.Open(f)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), pattern) {
			line := scanner.Text()
			if slices.Contains(flags, "-n") {
				line = fmt.Sprintf("%d:%s", lineNum+1, line)
			}
			if slices.Contains(flags, "-l") && !slices.Contains(result, f) {
				result = append(result, f)
			} else {
				result = append(result, line)
			}
		}
		lineNum++
	}
	return result
}
