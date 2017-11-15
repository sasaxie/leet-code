package reverse_integer

import (
	"math"
	"testing"
)

func TestReverse(t *testing.T) {
	if reverse(123) != 321 {
		t.Error("error: 123 reverse not equal 321")
	} else {
		t.Log("success: 123 reverse equal 321")
	}

	if reverse(120) != 21 {
		t.Error("error: 120 reverse not equal 21")
	} else {
		t.Log("success: 120 reverse equal 21")
	}

	if reverse(math.MaxInt32+1) != 0 {
		t.Errorf("error: %d reverse not equal 0", math.MaxInt32+1)
	} else {
		t.Logf("success: %d reverse equal 0", math.MaxInt32+1)
	}

	if reverse(math.MinInt32-1) != 0 {
		t.Errorf("error: %d reverse not equal 0", math.MinInt32-1)
	} else {
		t.Logf("success: %d reverse equal 0", math.MinInt32-1)
	}
}
