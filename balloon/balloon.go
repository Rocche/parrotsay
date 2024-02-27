// Thank you Flavio for the amazing tutorial "Building a CLI command with Go: cowsay"
// https://flaviocopes.com/go-tutorial-cowsay/
// Please go check it out!
package balloon

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// builds the balloon containing the text
func buildBalloon(lines []string, maxwidth int) string {
	var ret []string

	top := "+" + strings.Repeat("-", maxwidth+2) + "+"
	bottom := "+" + strings.Repeat("-", maxwidth+2) + "+"

	ret = append(ret, top)

	for _, line := range lines {
		s := fmt.Sprintf("| %s |", line)
		ret = append(ret, s)
	}

	ret = append(ret, bottom)
	ret = append(ret, strings.Repeat(" ", len(bottom)/2-1)+"\\")
	ret = append(ret, strings.Repeat(" ", len(bottom)/2)+"\\")
	return strings.Join(ret, "\n")
}

// converts tabs to spaces
func tabsToSpaces(lines []string) []string {
	var ret []string
	for _, l := range lines {
		l = strings.Replace(l, "\t", "    ", -1)
		ret = append(ret, l)
	}
	return ret
}

// calculate the width of the longest line
func calculateMaxWidth(lines []string) int {
	w := 0
	for _, l := range lines {
		len := utf8.RuneCountInString(l)
		if len > w {
			w = len
		}
	}
	return w
}

// Given the max width, it normalizes all the lines to match the max width,
// eventually filling the end of the line with spaces
func normalizeStringsLength(lines []string, maxwidth int) []string {
	var ret []string
	for _, l := range lines {
		nSpaces := maxwidth - utf8.RuneCountInString(l)
		s := l + strings.Repeat(" ", nSpaces)
		ret = append(ret, s)
	}
	return ret
}

// PrintBalloon prints the balloon containing the text
func PrintBalloon(lines []string) {
	noTabbedLines := tabsToSpaces(lines)
	maxwidth := calculateMaxWidth(noTabbedLines)
	messages := normalizeStringsLength(noTabbedLines, maxwidth)
	balloon := buildBalloon(messages, maxwidth)
	fmt.Println(balloon)
}
