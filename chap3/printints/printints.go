package printints

import (
	"bytes"
	"fmt"
)

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')

	for i, val := range values {
		fmt.Fprintf(&buf, "%d", val)
		if i < len(values)-1 {
			buf.WriteString(", ")
		}
	}

	buf.WriteByte(']')
	return buf.String()
}
