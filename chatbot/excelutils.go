package chatbotada

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	chatbotpb "github.com/zhs007/chatbot/chatbotpb"
)

// ExcelCellType - excel cell type
type ExcelCellType int32

const (
	// CellInvalid - invalid
	CellInvalid ExcelCellType = -1
	// CellString - string
	CellString ExcelCellType = 0
	// CellInt32 - int32
	CellInt32 ExcelCellType = 1
	// CellInt64 - int64
	CellInt64 ExcelCellType = 2
	// CellFloat32 - float32
	CellFloat32 ExcelCellType = 3
	// CellFloat64 - float64
	CellFloat64 ExcelCellType = 4
	// CellPercentage - percentage
	CellPercentage ExcelCellType = 5
	// CellNull - null
	CellNull ExcelCellType = 6
)

// ExcelColumnType - excel column type
type ExcelColumnType int32

const (
	// ColumnInfo - info
	ColumnInfo ExcelColumnType = 0
	// ColumnNumberPrimaryKey - number primary key
	ColumnNumberPrimaryKey ExcelColumnType = 1
	// ColumnPrimaryKey - primary key
	ColumnPrimaryKey ExcelColumnType = 2
	// ColumnDataTime - data time
	ColumnDataTime ExcelColumnType = 3
	// ColumnNumber - number
	ColumnNumber ExcelColumnType = 4
	// ColumnPercentage - percentage
	ColumnPercentage ExcelColumnType = 5
	// ColumnIgnorePrimaryKey - ignore primary key
	ColumnIgnorePrimaryKey ExcelColumnType = 6
	// ColumnCategory - category
	ColumnCategory ExcelColumnType = 7
	// ColumnInt - int
	ColumnInt ExcelColumnType = 8
	// ColumnTimestamp - timestamp
	ColumnTimestamp ExcelColumnType = 9
	// ColumnTimestampMs - timestampms
	ColumnTimestampMs ExcelColumnType = 10
	// ColumnNull - null
	ColumnNull ExcelColumnType = 11
	// ColumnTreeCategory - tree category
	ColumnTreeCategory ExcelColumnType = 12
	// ColumnIgnore - ignore
	ColumnIgnore ExcelColumnType = 13
	// ColumnMultiCategories - multiple categories
	ColumnMultiCategories ExcelColumnType = 14
)

// ExcelColumnTypeObj - excel column type object
type ExcelColumnTypeObj struct {
	Name      string
	Type      ExcelColumnType
	Separator string
}

var lstColumnString = []string{
	"Info",
	"PrimaryKey",
	"NumberPrimaryKey",
	"DataTime",
	"Number",
	"Percentage",
	"IgnorePrimaryKey",
	"Category",
	"Int",
	"Timestamp",
	"TimestampMs",
	"Null",
	"TreeCategory",
	"Ignore",
	"MultiCategories",
}

// ExcelData - excel data
type ExcelData struct {
	StartX       int
	StartY       int
	Columns      []ExcelColumnTypeObj
	ColumnsAuto  []ExcelColumnType
	Arr          [][]string
	CurSheetName string
}

func isFloat(str string) bool {
	const strFloat = "0123456789.%"
	// str = strings.TrimSpace(str)
	haspt := false
	for i, v := range str {
		if strings.IndexRune(strFloat, v) < 0 {
			return false
		}

		if v == '.' {
			if haspt {
				return false
			}

			haspt = true
		}

		if v == '%' {
			if i != len([]rune(str))-1 {
				return false
			}
		}
	}

	return true
}

func isPercentage(str string) bool {
	const strFloat = "0123456789.%"
	// str = strings.TrimSpace(str)
	haspt := false
	for i, v := range str {
		if strings.IndexRune(strFloat, v) < 0 {
			return false
		}

		if v == '.' {
			if haspt {
				return false
			}

			haspt = true
		}

		if v == '%' {
			if i != len([]rune(str))-1 {
				return false
			}

			return true
		}
	}

	return false
}

func isTimestamp(str string) bool {
	i64, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return false
	}
	if i64 > 100000000 && i64 < 10000000000 {
		return true
	}

	return false
}

func isTimestampMs(str string) bool {
	i64, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return false
	}
	if i64 > 100000000000 && i64 < 10000000000000 {
		return true
	}

	return false
}

func isInt(str string) bool {
	const strInt = "0123456789,"
	for _, v := range str {
		if strings.IndexRune(strInt, v) < 0 {
			return false
		}
	}

	return true
}

func isDataTime(str string) bool {
	_, err := time.Parse(time.RFC3339, str)
	if err == nil {
		return true
	}

	_, err = time.Parse("2006-01-02", str)
	if err == nil {
		return true
	}

	_, err = time.Parse("2006-01-02 15:04:05", str)
	if err == nil {
		return true
	}

	_, err = time.Parse("02/01/2006 15:04:05", str)
	if err == nil {
		return true
	}

	_, err = time.Parse("02/01/2006", str)
	if err == nil {
		return true
	}

	return false
}

// AnalysisCell - analysis cell, exclude rows with y <= sy
func AnalysisCell(arr [][]string, x int, sy int) ExcelCellType {
	if x >= 0 && x < len(arr[0]) {
		isneedfloat := false

		// 如果至少一行是float，则整列都应该是float
		// 如果至少一行不是数字，则整列都不应该是数字
		// 暂时不区分32位和64位，默认为64位

		isallnull := true
		for y := sy + 1; y < len(arr); y++ {
			cr := strings.TrimSpace(arr[y][x])
			if cr == "" {
				if y == len(arr)-1 && isallnull {
					return CellNull
				}

				continue
			}

			isallnull = false

			isf := isFloat(cr)
			if !isf {
				return CellString
			}

			if !isneedfloat {
				isi := isInt(cr)
				if !isi {
					isneedfloat = true
				}
			}
		}

		if isneedfloat {
			return CellFloat64
		}

		return CellInt64
	}

	return CellInvalid
}

func isHan(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}

	return false
}

func isColumnInfo(str string) bool {
	// 如果有控制符号，算info
	for _, r := range str {
		if unicode.IsControl(r) {
			return true
		}
	}

	// 如果没有标点符号，且没有空格，不算info
	nospaceandpunct := true
	for _, r := range str {
		if unicode.IsSpace(r) || unicode.IsPunct(r) {
			nospaceandpunct = false

			break
		}
	}

	if nospaceandpunct {
		return false
	}

	// 如果是汉字，超过8个汉字的算info
	if isHan(str) {
		nums := 0
		for _, r := range str {
			if unicode.Is(unicode.Scripts["Han"], r) {
				nums++
			}
		}

		if nums > 8 {
			return true
		}
	} else {
		// 如果不包含汉字，超过8个单词的，算info
		words := strings.Fields(str)
		if len(words) > 8 {
			return true
		}
	}

	// 如果长度超过32的，算info
	lstr := []rune(str)
	if len(lstr) > 32 {
		return true
	}

	return false
}

// hasDuplicationColumn - is there no duplication?
func hasDuplicationColumn(arr [][]string, x int, cv string, cy int) bool {
	for y := 1; y < len(arr); y++ {
		if y != cy && arr[y][x] == cv {
			return true
		}
	}

	return false
}

// HasDuplication - is there no duplication?
func HasDuplication(arr [][]string, x int) bool {
	for y := 1; y < len(arr); y++ {
		if hasDuplicationColumn(arr, x, arr[y][x], y) {
			return true
		}
	}

	return false
}

// AnalysisColumn - analysis column, exclude rows with y == 0
func AnalysisColumn(arr [][]string, x int, sy int) ExcelColumnType {
	ct := AnalysisCell(arr, x, sy)
	if ct == CellString {
		isinfo := false
		for y := 1; y < len(arr); y++ {
			if isColumnInfo(arr[y][x]) {
				isinfo = true

				break
			}
		}

		if isinfo {
			return ColumnInfo
		}

		isdatatime := true
		for y := 1; y < len(arr); y++ {
			if !isDataTime(arr[y][x]) {
				isdatatime = false

				break
			}
		}

		if isdatatime {
			return ColumnDataTime
		}

		if !HasDuplication(arr, x) {
			return ColumnPrimaryKey
		}

		return ColumnCategory
	} else if ct == CellFloat64 || ct == CellFloat32 || ct == CellPercentage {
		if ct == CellPercentage {
			return ColumnPercentage
		}

		for y := 1; y < len(arr); y++ {
			if !isPercentage(arr[y][x]) {
				return ColumnNumber
			}
		}

		return ColumnPercentage
	} else if ct == CellInt32 || ct == CellInt64 {
		istimestamp := true
		for y := 1; y < len(arr); y++ {
			if !isTimestamp(arr[y][x]) {
				istimestamp = false

				break
			}
		}

		if istimestamp {
			return ColumnTimestamp
		}

		istimestampms := true
		for y := 1; y < len(arr); y++ {
			if !isTimestampMs(arr[y][x]) {
				istimestampms = false

				break
			}
		}

		if istimestampms {
			return ColumnTimestampMs
		}

		if !HasDuplication(arr, x) {
			return ColumnNumberPrimaryKey
		}

		return ColumnInt
	}

	return ColumnInfo
}

// AnalysisColumnsType - analysis column type
func AnalysisColumnsType(arr [][]string, sx int, sy int) []ExcelColumnType {
	var lst []ExcelColumnType

	for x := sx; x < len(arr[0]); x++ {
		lst = append(lst, AnalysisColumn(arr, x, sy))
	}

	return lst
}

// ProcHead - process head
func ProcHead(arr [][]string, sx int, sy int) [][]string {
	for i, v := range arr[sy] {
		if i >= sx {
			if v == "" {
				arr[sy][i] = fmt.Sprintf("__column%v__", i)
			}
		}
	}

	return arr
}

// ExcelColumnType2String - ExcelColumnType -> string
func ExcelColumnType2String(ect ExcelColumnType) string {
	if ect >= 0 && int(ect) < len(lstColumnString) {
		return lstColumnString[ect]
	}

	return "invalid"
}

func isEmptyRow(arr [][]string, x int) bool {
	if x >= 0 && x < len(arr[0]) {
		for y := 0; y < len(arr); y++ {
			cr := strings.TrimSpace(arr[y][x])
			if cr != "" {
				return false
			}
		}

		return true
	}

	return true
}

func isEmptyColumn(arr [][]string, y int) bool {
	if y >= 0 && y < len(arr) {
		for x := 0; x < len(arr[0]); x++ {
			cr := strings.TrimSpace(arr[y][x])
			if cr != "" {
				return false
			}
		}

		return true
	}

	return true
}

// GetStartXY - get start x & y
func GetStartXY(arr [][]string) (int, int) {
	cx := 0
	for x := 0; x < len(arr[0]); x++ {
		if !isEmptyRow(arr, x) {
			cx = x

			break
		}
	}

	cy := 0
	for y := 0; y < len(arr); y++ {
		if !isEmptyColumn(arr, y) {
			cy = y

			break
		}
	}

	return cx, cy
}

// AnalysisColumnsTypeWithComments - analysis ColumnsType with comments
func AnalysisColumnsTypeWithComments(arr [][]string, sx int, sy int,
	mapComments map[string]string) []ExcelColumnTypeObj {

	var lst []ExcelColumnTypeObj

	for tx := sx; tx < len(arr[0]); tx++ {
		cn, err := excelize.CoordinatesToCellName(tx+1, sy+1)
		if err == nil {
			// fmt.Printf("%v\n", cn)

			d, isok := mapComments[cn]
			if isok {
				d = strings.TrimSpace(d)

				if d == "PrimaryKey" {
					lst = append(lst, ExcelColumnTypeObj{
						Name: arr[sy][tx],
						Type: ColumnPrimaryKey,
					})

					continue
				} else if strings.Contains(d, "TreeCategory") {
					ca := strings.Split(d, " ")
					if len(ca) > 1 {
						lst = append(lst, ExcelColumnTypeObj{
							Name:      arr[sy][tx],
							Type:      ColumnTreeCategory,
							Separator: strings.TrimSpace(ca[1]),
						})

						continue
					}
				} else if strings.Contains(d, "MultiCategories") {
					ca := strings.Split(d, " ")
					if len(ca) > 1 {
						lst = append(lst, ExcelColumnTypeObj{
							Name:      arr[sy][tx],
							Type:      ColumnMultiCategories,
							Separator: strings.TrimSpace(ca[1]),
						})

						continue
					}
				} else if d == "Category" {
					lst = append(lst, ExcelColumnTypeObj{
						Name: arr[sy][tx],
						Type: ColumnCategory,
					})

					continue
				} else if d == "Ignore" {
					lst = append(lst, ExcelColumnTypeObj{
						Name: arr[sy][tx],
						Type: ColumnIgnore,
					})

					continue
				} else if d == "Timestamp" {
					lst = append(lst, ExcelColumnTypeObj{
						Name: arr[sy][tx],
						Type: ColumnTimestamp,
					})

					continue
				} else if d == "TimestampMs" {
					lst = append(lst, ExcelColumnTypeObj{
						Name: arr[sy][tx],
						Type: ColumnTimestampMs,
					})

					continue
				}
			}
		}

		lst = append(lst, ExcelColumnTypeObj{
			Name: arr[sy][tx],
			Type: ColumnNull,
		})
	}

	return lst
}

// GetComments - get comments with sheetName
func GetComments(f *excelize.File, sheetName string) map[string]string {
	mc := f.GetComments()
	cursheetcomments, isok := mc[sheetName]
	if !isok {
		return nil
	}

	mapComments := make(map[string]string)
	for _, v := range cursheetcomments {
		mapComments[v.Ref] = v.Text
	}

	return mapComments
}

// ProcExcelMsg - analysis chatmsg to ExcelData
func ProcExcelMsg(chat *chatbotpb.ChatMsg) (*ExcelData, error) {
	r := bytes.NewReader(chat.File.FileData)
	f, err := excelize.OpenReader(r)
	if err != nil {
		return nil, err
	}

	cs := f.GetActiveSheetIndex()
	curSheet := f.GetSheetName(cs)

	arr, err := f.GetRows(curSheet)
	if err != nil {
		return nil, err
	}

	sx, sy := GetStartXY(arr)

	arr = ProcHead(arr, sx, sy)

	mapComments := GetComments(f, curSheet)

	lstct := AnalysisColumnsType(arr, sx, sy)
	lstctobj := AnalysisColumnsTypeWithComments(arr, sx, sy, mapComments)

	for i, v := range lstct {
		if lstctobj[i].Type == ColumnNull {
			lstctobj[i].Type = v
		}
	}

	ed := &ExcelData{
		StartX:       sx,
		StartY:       sy,
		Arr:          arr,
		ColumnsAuto:  lstct,
		Columns:      lstctobj,
		CurSheetName: curSheet,
	}

	return ed, nil
}
