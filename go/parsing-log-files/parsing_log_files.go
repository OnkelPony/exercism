package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	var re = regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[*~=-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	var result int
	re := regexp.MustCompile(`\".*(?i)password.*\"`)
	for _, line := range lines {
		if re.MatchString(line) {
			result++
		}
	}
	return result
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d*`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\w+)`)
	for i, line := range lines {
		match := (re.FindStringSubmatch(line))
		if len(match) > 1 {
			lines[i] = "[USR] " + match[1] + " " + line
		}
	}
	return lines
}
