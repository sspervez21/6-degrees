package sixdegrees

import "strings"

func NormalizeName(name string) string {
	return strings.Replace(name, " ", "+", -1)
}

func DenormalizeName(name string) string {
	return strings.Replace(name, "+", " ", -1)
}
