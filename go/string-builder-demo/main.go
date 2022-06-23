package main

import (
	"fmt"
	"strings"
)

func main() {
	str := join("Hello", " ", "World")
	fmt.Println(str)
}

func join(strs ...string) string {
	var ret string
	for _, str := range strs {
		ret += str
	}
	return ret
}

func joinWithStringBuilder(strs ...string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(str)
	}
	return sb.String()
}

func joinRunes(runes ...rune) string {
	var sb strings.Builder
	for _, r := range runes {
		sb.WriteRune(r)
	}
	return sb.String()
}

func joinedAndReverse(strs ...string) (string, string) {
	var sb strings.Builder
	for _, v := range strs {
		sb.WriteString(v)
	}
	joined := sb.String()
	sb.Reset()
	for i := len(strs) - 1; i >= 0; i-- {
		sb.WriteString(strs[i])
	}
	return joined, sb.String()
}
