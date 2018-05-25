package util

import (
	"bytes"
)

// ConcatStrings .
func ConcatStrings(strings ...string) string {
	var buffer bytes.Buffer
	for _, item := range strings {
		buffer.WriteString(item)
	}
	return buffer.String()
}
