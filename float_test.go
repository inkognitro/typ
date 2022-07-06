package typ

import (
	"testing"
)

func TestFloatsShouldBeEqualByPrecisionOfTwo(t *testing.T) {
	t1 := NewFloat(5.555).SetPrecision(2)
	t2 := NewFloat(5.5556).SetPrecision(2)
	if !t1.Equals(t2) {
		t.Errorf("floats are not equal")
	}
}

func TestFloatsShouldNotBeEqualByPrecisionOfThree(t *testing.T) {
	t1 := NewFloat(5.555).SetPrecision(2)
	t2 := NewFloat(5.5556).SetPrecision(3)
	if t1.Equals(t2) {
		t.Errorf("floats are equal")
	}
}
