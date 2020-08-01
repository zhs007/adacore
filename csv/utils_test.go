package csv

import (
	"strings"
	"testing"
)

func Test_formatString(t *testing.T) {
	in := []string{
		"abc",
		"a,bc",
		"\"",
		"\"abc\"",
		"\"a,bc\"",
	}

	out := []string{
		"abc",
		"\"a,bc\"",
		"\"\"\"\"",
		"\"\"\"abc\"\"\"",
		"\"\"\"a,bc\"\"\"",
	}

	for i, v := range in {
		r := formatString(v)

		if r != out[i] {
			t.Fatalf("Test_formatString formatString %v %v %v", i, r, out[i])
		}
	}

	t.Logf("Test_formatString OK")
}

func Test_Save2Csv(t *testing.T) {
	type lineData struct {
		data []string
	}

	d := []lineData{
		{data: []string{"l1 d1", "l1 d2", "l1,d3"}},
	}

	Save2Csv("../unittest/csv001.csv", []string{"head1", "head2", "head,3"}, len(d),
		func(i int, member string) (string, error) {
			if strings.Index(member, "1") >= 0 {
				return d[i].data[0], nil
			} else if strings.Index(member, "2") >= 0 {
				return d[i].data[1], nil
			} else if strings.Index(member, "3") >= 0 {
				return d[i].data[2], nil
			}

			return "", nil
		})

	t.Logf("Test_Save2Csv OK")
}
