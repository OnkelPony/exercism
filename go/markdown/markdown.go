package markdown

// implementation to refactor

import (
	"fmt"
	"strings"
)

// Render translates markdown to HTML
func Render(markdown string) string {
	headerLevel := 0
	markdown = strings.Replace(markdown, "__", "<strong>", 1)
	markdown = strings.Replace(markdown, "__", "</strong>", 1)
	markdown = strings.Replace(markdown, "_", "<em>", 1)
	markdown = strings.Replace(markdown, "_", "</em>", 1)
// pos moved to for cycle
	list := 0
	listOpened := false
	html := ""
	he := false
	for pos := 0; pos < len(markdown);{
		char := markdown[pos]
		if char == '#' {
			for char == '#' {
				headerLevel++
				pos++
				char = markdown[pos]
			}
			if headerLevel == 7 {
				html += fmt.Sprintf("<p>%s ", strings.Repeat("#", headerLevel))
			} else if he {
				html += "# "
				headerLevel--
			} else {
				html += fmt.Sprintf("<h%d>", headerLevel)
			}
			pos++
			continue
		}
		he = true
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
			pos += 2
			continue
		}
		if char == '\n' {
			if listOpened && strings.LastIndex(markdown, "\n") == pos && strings.LastIndex(markdown, "\n") > strings.LastIndex(markdown, "*") {
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
			pos++
			continue
		}
		html += string(char)
		pos++
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
