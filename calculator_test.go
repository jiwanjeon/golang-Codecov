package calculator

import "testing"

func TestAdd(t *testing.T) {
	if got, want := Add(2, 3), 3; got != want {
		t.Errorf("add method produced wrong result. expected: %d, got: %d", want, got)
	}
}
