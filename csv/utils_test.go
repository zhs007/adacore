package csv

import (
	"testing"
)

func Test_formatString(t *testing.T) {
	in := []string{
		"abc",
		"a,bc",
	}

	out := []string{
		"abc",
		"\"a,bc\"",
	}

	for i, v := range in {
		r := formatString(v)

		if r != out[i] {
			t.Fatalf("Test_formatString formatString %v %v %v", i, r, out[i])
		}
	}

	t.Logf("Test_formatString OK")
}
