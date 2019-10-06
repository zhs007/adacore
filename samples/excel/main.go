package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	chatbotada "github.com/zhs007/adacore/chatbot"
)

func testExcel(fn string) error {
	f, err := excelize.OpenFile(fn)
	if err != nil {
		return err
	}

	csi := f.GetActiveSheetIndex()
	csn := f.GetSheetName(csi)

	rows, err := f.GetRows(csn)
	if err != nil {
		return err
	}

	fmt.Printf("w - %v h - %v\n", len(rows[0]), len(rows))

	cx, cy := chatbotada.GetStartXY(rows)

	fmt.Printf("cx - %v cy - %v", cx, cy)

	return nil
}

func main() {
	testExcel("./excel001.xlsx")
}
