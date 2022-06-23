package main

import (
	"fmt"
	"strings"
	"testing"
)

var stringConcatTestTable = []struct {
	text []string
}{
	{text: []string{"Hello", " ", "From", " ", "Mars"}},
	{text: []string{"Hello", " ", "From", " ", "Venus"}},
	{text: []string{"Hello", " ", "From", " ", "Earth"}},
	{text: []string{"Hello", " ", "From", " ", "Saturn"}},
}

func BenchmarkSimpleJoin(b *testing.B) {
	for _, v := range stringConcatTestTable {
		b.Run(fmt.Sprintf("input_%s", v.text), func(b *testing.B) {
			_ = join(v.text...)

		})
	}
}

func BenchmarkJoinWithStringBuilder(b *testing.B) {
	fmt.Println("Executing benchmark")
	for _, v := range stringConcatTestTable {
		b.Run(fmt.Sprintf("input_%s", v.text), func(b *testing.B) {
			_ = joinWithStringBuilder(v.text...)
		})
	}

}

func TestJoinRunes(t *testing.T) {
	s := "golang"
	s_rune := []rune(s)
	result := joinRunes(s_rune...)
	fmt.Println(s_rune)

	if !strings.EqualFold(result, s) {
		t.Errorf("Expect: %s, got: %s", s, result)
	}
}

func TestJoinedAndReverse(t *testing.T) {
	result, reversed := joinedAndReverse("Hello", " ", "from", " ", "Mars")
	if !strings.EqualFold(result, "Hello from Mars") {
		t.Errorf("Expercted: 'Hello from Mars', got: %s", result)
	}
	if !strings.EqualFold(reversed, "Mars from Hello") {
		t.Errorf("Expercted: 'Mars from Hello', got: %s", reversed)
	}
}
