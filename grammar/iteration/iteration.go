package iteration

import "strings"

func Repeat(s string) string {
	var repeated strings.Builder
	for i := 0; i < 3; i++ {
		repeated.WriteString(s)
	}
	return repeated.String()
}
