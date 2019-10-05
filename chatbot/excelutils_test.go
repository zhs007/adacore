package chatbotada

import (
	"testing"
)

func Test_isDataTime(t *testing.T) {
	lstok := []string{
		"2019-10-01 01:02:59",
		"2019-10-01",
		"01/10/2019 01:02:59",
		"01/10/2019",
	}

	lstfail := []string{
		"2019-09-31",
		"2019-10-01 24:00:00",
		"13/13/2019",
	}

	for _, v := range lstok {
		cr := isDataTime(v)

		if !cr {
			t.Fatalf("Test_isDataTime invalid ok %v", v)

			return
		}
	}

	for _, v := range lstfail {
		cr := isDataTime(v)

		if cr {
			t.Fatalf("Test_isDataTime invalid fail %v", v)

			return
		}
	}

	t.Logf("Test_isDataTime OK")
}
