package problem002

import "testing"

func TestProblem002_40(t *testing.T) {
	sum := Problem002(40)
	expected := 44
	if sum != expected {
		t.Error("For Problem2 expected", expected, "got", sum)
	}
}
func TestProblem002_4000000(t *testing.T) {
	sum := Problem002(4000000)
	expected := 4613732
	if sum != expected {
		t.Error("For Problem2 expected", expected, "got", sum)
	}
}
