package text

import (
	"regexp"
	"strings"
)

func CleanEmail(input string) string {
	text := strings.ToLower(input)
	re := regexp.MustCompile(`[^a-z0-9_.@]`)
	return re.ReplaceAllString(text, "")
}