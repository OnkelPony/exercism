package grep

import (
	"bufio"
	"os"
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
	file, _ := os.Open(f)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), pattern) {
			result = append(result, scanner.Text())
		}
	}
	return result
}

