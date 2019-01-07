package fizzbuzz_tdd_go

import "testing"

func TestInput_1_Plus_1_ShouldBe_2(t *testing.T) {
	fb := plusOneAndOne(1, 1)
	expected := "2"
	if fb != expected {
		t.Errorf("Actual = %s, expected = %s", fb, expected)
	}
}
