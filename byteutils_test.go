package byteutils

import "testing"

func SanityCheck(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Fatalf("1 + 2 != 3")
	}
}
