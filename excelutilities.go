package excel_utils

import (
	excelize "github.com/360EntSecGroup-Skylar/excelize"
)

var line, col, auforFilterStartcol, auforFilterStartrow int
var fexcel *excelize.File
var sheet string

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
func NewFile() {
	fexcel = excelize.NewFile()
	line = 1
	col = 1
	sheet = "Sheet1"
}
func NextLine() {
	line++
}
func NextCol() {
	col++
}
func ResetCol() {
	col = 1
}

func WriteColumnsHeaderln(data []string) {
	for _, v := range data {
		SetTableHeader()
		WiteCell(v)
		col++
	}
	col = 1
	line++
}
func WriteColumnsHeaderRotln(data []string) {
	for _, v := range data {
		SetTableHeader()
		WiteCell(v)
		col++
	}
	col = 1
	line++
}


func WriteColumns(data []string) {
	for _, v := range data {
		WiteCell(v)
		col++
	}
	col = 1
}

func WriteColumnsln(data []string) {
	WriteColumns(data)
	line++
}

func WiteCell(msg interface{}) {
	axis, err := excelize.CoordinatesToCellName(col, line)
	Check(err)
	err = fexcel.SetCellValue(sheet, axis, msg)
	Check(err)
}
func WiteCellln(msg interface{}) {
	WiteCell(msg)
	line++
}
func WiteCellnc(msg interface{}) {
	WiteCell(msg)
	col++
}

func SetCellStyleRotateXY(q, v int) {
	axis, err := excelize.CoordinatesToCellName(q, v)
	Check(err)
	style, err := fexcel.NewStyle(`{"alignment":{"text_rotation":90,"horizontal":"center"},"fill":{"type":"pattern","color":["#E0EBF5"],"pattern":1}}`)
	Check(err)
	err = fexcel.SetCellStyle(sheet, axis, axis, style)
	Check(err)
}
func SetCellStyleCenter() {
	axis, err := excelize.CoordinatesToCellName(col, line)
	Check(err)
	//var style excelize.Style
	style, err := fexcel.NewStyle(`{"alignment":{"horizontal":"center"}}`)
	Check(err)
	err = fexcel.SetCellStyle(sheet, axis, axis, style)
	Check(err)
}


func SetCellStyleRotate() {
	SetCellStyleRotateXY(col, line)
}
func SetCellStyleRotateN(count int) {
	var i, j int
	i = col
	j = line
	for n := 0; n < count; n++ {
		SetCellStyleRotateXY(i, j)
		i++
	}
}
func SetCellFontHeader() {
	axis, err := excelize.CoordinatesToCellName(col, line)
	Check(err)
	//var style excelize.Style
	style, err := fexcel.NewStyle(`{"font":{"bold":true,"family":"Times New Roman","size":24,"color":"#777777"}}`)
	Check(err)
	err = fexcel.SetCellStyle(sheet, axis, axis, style)
	Check(err)
}
func SetCellFontHeader2() {
	axis, err := excelize.CoordinatesToCellName(col, line)
	Check(err)
	//var style excelize.Style
	style, err := fexcel.NewStyle(`{"font":{"bold":true,"family":"Times New Roman","size":16,"color":"#777777"}}`)
	Check(err)
	err = fexcel.SetCellStyle(sheet, axis, axis, style)
	Check(err)
}
func SetTableHeader() {
	cellname, err := excelize.CoordinatesToCellName(col, line)
	Check(err)
	style, err := fexcel.NewStyle(`{"alignment":{"horizontal":"center","vertical":"center"},"fill":{"type":"pattern","color":["#E0EBF5"],"pattern":1}}`)
	Check(err)
	err = fexcel.SetCellStyle(sheet, cellname, cellname, style)
	Check(err)
}
func SetTableHeaderRot() {
	cellname, err := excelize.CoordinatesToCellName(col, line)
	Check(err)
	style, err := fexcel.NewStyle(`{"alignment":{"text_rotation":90,"horizontal":"center","vertical":"center"},"fill":{"type":"pattern","color":["#E0EBF5"],"pattern":1}}`)
	Check(err)
	err = fexcel.SetCellStyle(sheet, cellname, cellname, style)
	Check(err)
}

func AutoFilterStart() {
	auforFilterStartcol = col
	auforFilterStartrow = line
}
func AutoFilterEnd() {
	axis, err := excelize.CoordinatesToCellName(auforFilterStartcol, auforFilterStartrow)
	Check(err)
	autoFilter(axis)
}

func autoFilter(uppperleft string) {
	rows, err := fexcel.GetRows(sheet)
	Check(err)
	nrows := len(rows)
	ncols := len(rows[nrows-1])
	axis, err := excelize.CoordinatesToCellName(ncols, nrows)
	err = fexcel.AutoFilter(sheet, uppperleft, axis, "")
}

func SetColWidth(startcol, endcol string, width float64) {
	err := fexcel.SetColWidth(sheet, startcol, endcol, width)
	Check(err)
}
func SaveAs(name string) {
	err := fexcel.SaveAs(name)
	Check(err)

}
