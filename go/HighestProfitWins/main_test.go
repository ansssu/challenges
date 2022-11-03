package main

import "testing"

func TestMinMax(t *testing.T) {
	got := MinMax([]int{5, 3, 3, 9, 10, 3, 2, 1})
	if got != [2]int{1, 10} {
		t.Error("comparison failed")
	}
}
