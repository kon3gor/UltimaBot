package util

import (
	"regexp"
	"strings"
)

var runesToEscape = []string{
	`\(`,
	`\)`,
	`\_`,
	`\.`,
	`\*`,
	`\[`,
	`\]`,
}

func EscapeFakeMarkdown(s string) string {
	re := escapeRegexp()
	ind := re.FindAllStringIndex(s, -1)
	beginInd := make([]int, 0, len(ind))
	for _, v := range ind {
		beginInd = append(beginInd, v[0])
	}
	b := strings.Builder{}
	curBegInd := 0
	// todo: this shit will break immediatly when I try to use markdown is msges
	for i, c := range s {
		switch {
		case curBegInd < len(beginInd) && i == beginInd[curBegInd] && isEscapeSymbol(c):
			b.WriteRune('\\')
			b.WriteRune(c)
		case curBegInd < len(beginInd) && i == beginInd[curBegInd]:
			b.WriteRune(c)
			b.WriteRune('\\')
		default:
			b.WriteRune(c)
			continue
		}
		curBegInd++
	}
	return b.String()
}

func isEscapeSymbol(r rune) bool {
	isLowerRune := r >= 'a' && r <= 'z'
	isUpperRune := r >= 'A' && r <= 'Z'
	isNumber := r >= '0' && r <= '9'
	isWhitespacee := r == ' '
	return !isLowerRune && !isUpperRune && !isNumber && !isWhitespacee
}

func escapeRegexp() *regexp.Regexp {
	builder := strings.Builder{}
	builder.WriteString(`(^|[a-zA-Z0-9\s])(`)
	for _, r := range runesToEscape[1:] {
		builder.WriteString(r)
		builder.WriteRune('|')
	}
	builder.WriteString(runesToEscape[0])
	builder.WriteString(`)([a-zA-Z0-9\s]|$)`)
	return regexp.MustCompile(builder.String())
}
