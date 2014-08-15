package problem003

import "testing"

func TestProblem003_small(t *testing.T) {
	largestPrimeFactor := Problem003(13195)
	expected := 29
	if largestPrimeFactor != expected {
		t.Error("For Problem3 expected", expected, "got", largestPrimeFactor)
	}
}

func TestProblem003_reallyBig(t *testing.T) {
	largestPrimeFactor := Problem003(600851475143)
	expected := 6857
	if largestPrimeFactor != expected {
		t.Error("For Problem3 expected", expected, "got", largestPrimeFactor)
	}
}
