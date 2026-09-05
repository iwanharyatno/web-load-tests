package services

import (
	"fmt"
	"strings"
)

func GenerateBibNumber(prefix string, padding int, increment int) string {
	return fmt.Sprintf("%s%0*d", prefix, padding, increment)
}

func PadBibNumber(prefix string, padding, number int) string {
	var sb strings.Builder
	sb.WriteString(prefix)
	for i := 0; i < padding-len(fmt.Sprintf("%d", number)); i++ {
		sb.WriteString("0")
	}
	sb.WriteString(fmt.Sprintf("%d", number))
	return sb.String()
}
