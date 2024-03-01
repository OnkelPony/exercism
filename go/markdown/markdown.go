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
	listOpened := false
	// string.Builder improves performance
	var html strings.Builder
	var headerLevel int
	for i := 0; i < len(markdown); i++ {
		char := markdown[i]
		//changed from if to switch
		switch {
		case char == '#' && i == 0:
			headerLevel = countHeaderLevel(markdown, &i)
			if headerLevel < 7 {
				html.WriteString(fmt.Sprintf("<h%d>", headerLevel))
			} else {
				html.WriteString(fmt.Sprintf("<p>%s ", strings.Repeat("#", headerLevel)))
			}
		// header level test is not necessary
		case char == '*' && strings.Contains(markdown, "\n"):
			// move to separate function
			handleLists(&html, &listOpened, char)
			i++
		case char == '\n':
			// simplified the test
			if listOpened {
				// move to separate function
				listOpened = handleListItems(&html, markdown, i)
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
	switch {
	case strings.Contains(html.String(), "<ul>") && !strings.Contains(html.String(), "</ul>"):
		html.WriteString("</li></ul>")
	case strings.Contains(html.String(), "<p>"):
		html.WriteString("</p>")
	case headerLevel > 0:
		html.WriteString(fmt.Sprintf("</h%d>", headerLevel))
	default:
		return "<p>" + html.String() + "</p>"
	}
	return html.String()
}

func handleListItems(html *strings.Builder, markdown string, i int)bool {
	html.WriteString("</li>")
	if strings.LastIndex(markdown, "\n") == i && i > strings.LastIndex(markdown, "*") {
		html.WriteString("</ul><p>")
	}
	return false
}

func handleLists(html *strings.Builder, listOpened *bool, char byte)  {
	if !strings.Contains(html.String(), "<ul>") {
		html.WriteString("<ul>")
	}
	if !*listOpened {
		html.WriteString("<li>")
		*listOpened = true
	} else {
		html.WriteString(string(char) + " ")
	}
}

func countHeaderLevel(markdown string, i *int) int {
	headerLevel := 0
	for markdown[*i] == '#' {
		headerLevel++
		*i++
	}
	return headerLevel
}

func shapeLetters(html *string) {
	*html = strings.Replace(*html, "__", "<strong>", 1)
	*html = strings.Replace(*html, "__", "</strong>", 1)
	*html = strings.Replace(*html, "_", "<em>", 1)
	*html = strings.Replace(*html, "_", "</em>", 1)
}
