package chatbotada

import (
	"testing"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
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

func Test_AnalysisColumnsType(t *testing.T) {
	f, err := excelize.OpenFile("../unittest/excel001.xlsx")
	if err != nil {
		t.Fatalf("Test_AnalysisColumnsType OpenFile %v", err)

		return
	}

	cs := f.GetActiveSheetIndex()
	curSheet := f.GetSheetName(cs)

	arr, err := f.GetRows(curSheet)
	if err != nil {
		t.Fatalf("Test_AnalysisColumnsType GetRows %v", err)

		return
	}

	if len(arr) != 14 {
		t.Fatalf("Test_AnalysisColumnsType GetRows %v", len(arr))

		return
	}

	if len(arr[0]) != 5 {
		t.Fatalf("Test_AnalysisColumnsType GetRows %v", len(arr[0]))

		return
	}

	sx, sy := GetStartXY(arr)

	if sx != 1 || sy != 1 {
		t.Fatalf("Test_AnalysisColumnsType GetStartXY %v, %v", sx, sy)

		return
	}

	arr = ProcHead(arr, sx, sy)
	lstct := AnalysisColumnsType(arr, sx, sy)

	if len(lstct) != 4 {
		t.Fatalf("Test_AnalysisColumnsType fail %v", lstct)

		return
	}

	results := []ExcelColumnType{
		ColumnInt,
		ColumnPrimaryKey,
		ColumnInfo,
		ColumnInfo,
	}

	for i, v := range results {
		if v != lstct[i] {
			t.Fatalf("Test_AnalysisColumnsType invalid %v-%v", i, lstct[i])

			return
		}
	}

	t.Logf("Test_AnalysisColumnsType OK")
}

func Test_AnalysisColumnsTypeWithComments(t *testing.T) {
	f, err := excelize.OpenFile("../unittest/excel001.xlsx")
	if err != nil {
		t.Fatalf("Test_AnalysisColumnsTypeWithComments OpenFile %v", err)

		return
	}

	cs := f.GetActiveSheetIndex()
	curSheet := f.GetSheetName(cs)

	arr, err := f.GetRows(curSheet)
	if err != nil {
		t.Fatalf("Test_AnalysisColumnsTypeWithComments GetRows %v", err)

		return
	}

	sx, sy := GetStartXY(arr)

	arr = ProcHead(arr, sx, sy)

	mc := GetComments(f, curSheet)

	lstct := AnalysisColumnsTypeWithComments(arr, sx, sy, mc)

	if len(lstct) != 4 {
		t.Fatalf("Test_AnalysisColumnsTypeWithComments fail %v", lstct)

		return
	}

	results := []ExcelColumnType{
		ColumnNull,
		ColumnPrimaryKey,
		ColumnNull,
		ColumnNull,
	}

	for i, v := range results {
		if v != lstct[i].Type {
			t.Fatalf("Test_AnalysisColumnsTypeWithComments invalid %v-%v (%+v)", i, lstct[i], mc)

			return
		}
	}

	t.Logf("Test_AnalysisColumnsTypeWithComments OK")
}
