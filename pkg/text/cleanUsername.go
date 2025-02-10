package text

import (
	"regexp"
	"strings"
)

func CleanUsername(input string) string {
	text := strings.ToLower(input)
	re := regexp.MustCompile(`[^a-zA-Z0-9_.]`)
	return re.ReplaceAllString(text, "")
}