package iteration

import "testing"

func TestRepeat(t *testing.T) {

	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected '%s' but got '%s'", expected, repeated)
	}

}