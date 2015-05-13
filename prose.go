package prose

import "strings"

func Wrap(s string, width int) string {
	var lines []string
	var line string
	words := strings.Split(s, " ")
	for _, w := range words {
		var n string
		if line == "" {
			n = w
		} else {
			n = line + " " + w
		}
		if len(n) <= width {
			line = n
		} else {
			lines = append(lines, line)
			line = w
		}
	}
	lines = append(lines, line)
	return strings.Join(lines, "\n")
}
