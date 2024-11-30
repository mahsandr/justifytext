package main

import (
	"fmt"
	"strings"
)

func PrintWords(words []string, maxLineLen int) []string {
	startLine := 0
	tempLen := maxLineLen

	lines := make([]string, 0)
	var lenCurrentWord int
	lenAllWords := len(words)
	i := 0
	for i < lenAllWords {
		lenCurrentWord = len(words[i])
		if tempLen-lenCurrentWord >= 0 {
			tempLen -= lenCurrentWord + 1
		}
		if i == lenAllWords-1 || tempLen-len(words[i+1]) < 0 {
			var line []string
			if i == lenAllWords-1 {
				line = words[startLine:]
			} else {
				line = words[startLine : i+1]
			}
			lines = append(lines, addSpace(line, tempLen+1, i == lenAllWords-1))
			startLine = i + 1
			tempLen = maxLineLen
		}
		i++
	}
	return lines
}

func addSpace(words []string, remianChar int, isLastLint bool) string {
	moreSpace := 0
	if len(words)-1 > 0 {
		moreSpace = remianChar / (len(words) - 1)
	}
	line := words[0]
	space := " "
	if !isLastLint {
		space += strings.Repeat(" ", moreSpace)
	}
	for i := 1; i < len(words); i++ {
		if i == len(words)-1 && remianChar%(len(words)-1) != 0 {
			space += strings.Repeat(" ", remianChar%(len(words)-1))
		}
		line += space + words[i]
	}
	return line
}

func main() {
	lines := PrintWords([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16)
	for _, line := range lines {
		fmt.Println(line)
	}
}
