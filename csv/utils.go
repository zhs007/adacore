package csv

import (
	"os"
	"strings"

	adacorebase "github.com/zhs007/adacore/base"
)

// FuncGetCellString - getCellString(i int, member string) (string, error)
type FuncGetCellString func(i int, member string) (string, error)

func formatString(str string) string {
	if strings.Index(str, ",") >= 0 {
		return adacorebase.AppendString("\"", str, "\"")
	}

	return str
}

// Save2Csv - save to a csv file
func Save2Csv(fn string, lsthead []string, nums int, funcGetCellString FuncGetCellString) error {
	f, err := os.Create(fn)
	if err != nil {
		return err
	}

	defer f.Close()

	for i, vh := range lsthead {
		f.WriteString(formatString(vh))
		if i < len(lsthead)-1 {
			f.WriteString(", ")
		}
	}

	f.WriteString("\r\n")

	for i := 0; i < nums; i++ {
		for hi, hv := range lsthead {
			cv, err := funcGetCellString(i, hv)
			if err != nil {
				return err
			}

			f.WriteString(formatString(cv))

			if hi < len(lsthead)-1 {
				f.WriteString(", ")
			}
		}

		f.WriteString("\r\n")
	}

	return nil
}
