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
	list := 0
	listOpened := false
	html := ""
	headerLevel := 0
	for i := 0; i < len(markdown); {
		char := markdown[i]
		if char == '#' && i == 0{
			for char == '#' {
				headerLevel++
				i++
				char = markdown[i]
			}
			if headerLevel > 6 {
				html += fmt.Sprintf("<p>%s ", strings.Repeat("#", headerLevel))
			} else {
				html += fmt.Sprintf("<h%d>", headerLevel)
			}
			i++
			continue
		}
		if char == '*' && headerLevel == 0 && strings.Contains(markdown, "\n") {
			if list == 0 {
				html += "<ul>"
			}
			list++
			if !listOpened {
				html += "<li>"
				listOpened = true
			} else {
				html += string(char) + " "
			}
			i += 2
			continue
		}
		if char == '\n' {
			if listOpened && strings.LastIndex(markdown, "\n") == i && strings.LastIndex(markdown, "\n") > strings.LastIndex(markdown, "*") {
				html += "</li></ul><p>"
				listOpened = false
				list = 0
			}
			if list > 0 && listOpened {
				html += "</li>"
				listOpened = false
			}
			if headerLevel > 0 {
				html += fmt.Sprintf("</h%d>", headerLevel)
				headerLevel = 0
			}
			i++
			continue
		}
		html += string(char)
		i++
		//removed break
	}
	switch {
	case headerLevel == 7:
		return html + "</p>"
	case headerLevel > 0:
		return html + fmt.Sprintf("</h%d>", headerLevel)
	}
	if list > 0 {
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
