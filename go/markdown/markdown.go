package markdown

// implementation to refactor

import (
	"fmt"
	"strings"
)

// Render translates markdown to HTML
func Render(markdown string) string {
	shapeLetters(&markdown)
	// pos moved to for cycle
	isList := false
	listOpened := false
	html := ""
	headerLevel := 0
	for i := 0; i < len(markdown); i++ {
		char := markdown[i]
		if char == '#' && i == 0 {
			for markdown[i] == '#' {
				headerLevel++
				i++
			}
			if headerLevel < 7 {
				html += fmt.Sprintf("<h%d>", headerLevel)
			} else {
				html += fmt.Sprintf("<p>%s ", strings.Repeat("#", headerLevel))
			}
			continue
		}
		// header level test is not necessary
		if char == '*' && strings.Contains(markdown, "\n") {
			if !strings.Contains(html, "<ul>") {
				html += "<ul>"
			}
			isList = true
			if !listOpened {
				html += "<li>"
				listOpened = true
			} else {
				html += string(char) + " "
			}
			i ++
			continue
		}
		if char == '\n' {
			// simplified the test
			if listOpened && strings.LastIndex(markdown, "\n") == i && i > strings.LastIndex(markdown, "*") {
				html += "</li></ul><p>"
				listOpened = false
				isList = false
			}
			if isList  && listOpened {
				html += "</li>"
				listOpened = false
			}
			if headerLevel > 0 {
				html += fmt.Sprintf("</h%d>", headerLevel)
				headerLevel = 0
			}
			continue
		}
		html += string(char)
		//removed break
	}
	switch {
	case headerLevel == 7:
		return html + "</p>"
	case headerLevel > 0:
		return html + fmt.Sprintf("</h%d>", headerLevel)
	}
	if isList  {
		return html + "</li></ul>"
	}
	if strings.Contains(html, "<p>") {
		return html + "</p>"
	}
	return "<p>" + html + "</p>"

}

func shapeLetters(markdown *string) {
	*markdown = strings.Replace(*markdown, "__", "<strong>", 1)
	*markdown = strings.Replace(*markdown, "__", "</strong>", 1)
	*markdown = strings.Replace(*markdown, "_", "<em>", 1)
	*markdown = strings.Replace(*markdown, "_", "</em>", 1)
}
