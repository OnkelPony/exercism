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
	// string.Builder improves performance
	var html strings.Builder
	headerLevel := 0
	for i := 0; i < len(markdown); i++ {
		char := markdown[i]
		//changed from if to switch
		switch {
		case char == '#' && i == 0:
			for markdown[i] == '#' {
				headerLevel++
				i++
			}
			if headerLevel < 7 {
				html.WriteString(fmt.Sprintf("<h%d>", headerLevel))
			} else {
				html.WriteString(fmt.Sprintf("<p>%s ", strings.Repeat("#", headerLevel)))
			}
		// header level test is not necessary
		case char == '*' && strings.Contains(markdown, "\n"):
			if !isList {
				html.WriteString("<ul>")
			}
			isList = true
			if !listOpened {
				html.WriteString("<li>")
				listOpened = true
			} else {
				html.WriteString(string(char) + " ")
			}
			i++
		case char == '\n':
			// simplified the test
			if listOpened {
				html.WriteString("</li>")
				if strings.LastIndex(markdown, "\n") == i && i > strings.LastIndex(markdown, "*") {
					html.WriteString("</ul><p>")
					isList = false
				}
				listOpened = false
			}
			if headerLevel > 0 {
				html.WriteString(fmt.Sprintf("</h%d>", headerLevel))
				headerLevel = 0
			}
		default:
			html.WriteByte(char)
			//removed break
		}
	}
	// shape letters later improves performance
	switch {
	case headerLevel == 7:
		 html.WriteString("</p>")
	case headerLevel > 0:
		 html.WriteString(fmt.Sprintf("</h%d>", headerLevel))
	case isList:
		 html.WriteString("</li></ul>")
	case strings.Contains(html.String(), "<p>"):
		 html.WriteString("</p>")
	default:
		return "<p>" + html.String() + "</p>"
	}
	return html.String()
}

func shapeLetters(html *string) {
	*html = strings.Replace(*html, "__", "<strong>", 1)
	*html = strings.Replace(*html, "__", "</strong>", 1)
	*html = strings.Replace(*html, "_", "<em>", 1)
	*html = strings.Replace(*html, "_", "</em>", 1)
}
