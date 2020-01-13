package excel

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

// SetRow - set row
func SetRow(f *excelize.File, sheet string, column int, row int, lst []string) error {
	for i, v := range lst {
		cn, err := excelize.CoordinatesToCellName(column+i, row)
		if err != nil {
			return err
		}

		f.SetCellValue(sheet, cn, v)
	}

	return nil
}

// FuncGetCellWithPB - getcellwithpb(i int, member string) (interface{}, error)
type FuncGetCellWithPB func(i int, member string) (interface{}, error)

// FuncGetHeadComment - getheadcomment(i int, head string) string
type FuncGetHeadComment func(i int, head string) string

// FuncIsIgnoreRow - isignorerow(i int) bool
type FuncIsIgnoreRow func(i int) bool

// SetSheet - set sheet
func SetSheet(f *excelize.File, sheet string, column int, row int, lsthead []string, nums int,
	funcGetHeadComment FuncGetHeadComment, funcGetCellWithPB FuncGetCellWithPB) error {

	for i, v := range lsthead {
		cn, err := excelize.CoordinatesToCellName(column+i, row)
		if err != nil {
			return err
		}

		f.SetCellValue(sheet, cn, v)

		comment := funcGetHeadComment(i, v)
		if comment != "" {
			f.AddComment(sheet, cn, fmt.Sprintf(`{"author":" ", "text":"%v"}`, comment))
		}
	}

	for i := 0; i < nums; i++ {
		for hi, hv := range lsthead {
			cv, err := funcGetCellWithPB(i, hv)
			if err != nil {
				return err
			}

			cn, err := excelize.CoordinatesToCellName(column+hi, row+i+1)
			if err != nil {
				return err
			}

			f.SetCellValue(sheet, cn, cv)
		}
	}

	return nil
}

// SetSheetEx - set sheet
func SetSheetEx(f *excelize.File, sheet string, column int, row int, lsthead []string, nums int,
	funcGetHeadComment FuncGetHeadComment, funcIsIgnoreRow FuncIsIgnoreRow, funcGetCellWithPB FuncGetCellWithPB) error {

	for i, v := range lsthead {
		cn, err := excelize.CoordinatesToCellName(column+i, row)
		if err != nil {
			return err
		}

		f.SetCellValue(sheet, cn, v)

		comment := funcGetHeadComment(i, v)
		if comment != "" {
			f.AddComment(sheet, cn, fmt.Sprintf(`{"author":" ", "text":"%v"}`, comment))
		}
	}

	ti := 1
	for i := 0; i < nums; i++ {
		if funcIsIgnoreRow(i) {
			continue
		}

		for hi, hv := range lsthead {
			cv, err := funcGetCellWithPB(i, hv)
			if err != nil {
				return err
			}

			cn, err := excelize.CoordinatesToCellName(column+hi, row+ti)
			if err != nil {
				return err
			}

			f.SetCellValue(sheet, cn, cv)
		}

		ti++
	}

	return nil
}
