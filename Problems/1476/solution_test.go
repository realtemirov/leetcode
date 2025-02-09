package problem1476_test

import (
	"testing"

	problem1476 "github.com/realtemirov/leetcode/Problems/1476"
)

func TestSolution(t *testing.T) {

	rectangle := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}

	obj := problem1476.Constructor(rectangle)

	// Test initial values
	if got := obj.GetValue(0, 0); got != 1 {
		t.Errorf("GetValue(0, 0) = %d; want 1", got)
	}
	if got := obj.GetValue(1, 2); got != 7 {
		t.Errorf("GetValue(1, 2) = %d; want 7", got)
	}
	if got := obj.GetValue(2, 3); got != 12 {
		t.Errorf("GetValue(2, 3) = %d; want 12", got)
	}

	// Test UpdateSubrectangle
	obj.UpdateSubrectangle(1, 1, 2, 2, 99)

	// After update, subrectangle from (1,1) to (2,2) should be 99
	if got := obj.GetValue(1, 1); got != 99 {
		t.Errorf("GetValue(1, 1) = %d; want 99", got)
	}
	if got := obj.GetValue(2, 2); got != 99 {
		t.Errorf("GetValue(2, 2) = %d; want 99", got)
	}
	if got := obj.GetValue(1, 2); got != 99 {
		t.Errorf("GetValue(1, 2) = %d; want 99", got)
	}
	if got := obj.GetValue(2, 1); got != 99 {
		t.Errorf("GetValue(2, 1) = %d; want 99", got)
	}

	// Test non-updated values
	if got := obj.GetValue(0, 0); got != 1 {
		t.Errorf("GetValue(0, 0) = %d; want 1", got)
	}
	if got := obj.GetValue(0, 3); got != 4 {
		t.Errorf("GetValue(0, 3) = %d; want 4", got)
	}
	if got := obj.GetValue(2, 3); got != 12 {
		t.Errorf("GetValue(2, 3) = %d; want 12", got)
	}

	// Test larger update range
	obj.UpdateSubrectangle(0, 0, 1, 1, 77)

	// After update, the rectangle (0,0) to (1,1) should be 77
	if got := obj.GetValue(0, 0); got != 77 {
		t.Errorf("GetValue(0, 0) = %d; want 77", got)
	}
	if got := obj.GetValue(1, 1); got != 77 {
		t.Errorf("GetValue(1, 1) = %d; want 77", got)
	}
}
