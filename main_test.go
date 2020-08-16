package solution

import "testing"

func TestAdd(t *testing.T) {
	get := add(1, 1)
	want := 2
	if get != want {
		t.Errorf("want %v but get %v", want, get)
	}
}
