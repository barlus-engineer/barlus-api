package cleantext

import "regexp"

func SpecialChar(input string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9_]`)
	return re.ReplaceAllString(input, "")
}