package utils

import "regexp"

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)

func ClearStr(str string) string {
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}
