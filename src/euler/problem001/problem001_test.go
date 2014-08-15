package problem001

import "testing"

func TestProblem001_10(t *testing.T) {
	sum := Problem001(10)
	expected := 23
	if sum != expected {
		t.Error("For Problem1 expected", expected, "got", sum)
	}
}

func TestProblem001_1000(t *testing.T) {
	sum := Problem001(1000)
	expected := 233168
	if sum != expected {
		t.Error("For Problem1 expected", expected, "got", sum)
	}
}
