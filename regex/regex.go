// Package regex provides usefull regular expressions.
package regex

import "regexp"

const (
	alphaPattern  = "^[a-zA-Z]+$"
	numberPattern = "^[0-9]+$"
)

var (
	Alpha  = regexp.MustCompile(alphaPattern)
	Number = regexp.MustCompile(numberPattern)
)
