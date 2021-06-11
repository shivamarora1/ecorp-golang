package custom_strings

import (
	"fmt"
)

func Len(s string) string {
	return fmt.Sprintf("%d", len(s))
}

func ReverseConcatenate(a, b string) string {
	return b + a
}
