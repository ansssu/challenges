package kata

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {

	currentResult := TwoSum([]int{1, 2, 3}, 4)
	if !reflect.DeepEqual(currentResult, [2]int{0, 2}) {
		t.Errorf("Expected: %v  got: %v ", []int{0, 2}, currentResult)
	}

}
