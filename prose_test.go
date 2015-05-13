package prose

import "testing"

func TestWrap(t *testing.T) {
	s := "Hello from Denver!"
	r := Wrap(s, 10)
	if r != "Hello from\nDenver!" {
		t.Error("\nGot: ", "\n"+r)
	}
}
