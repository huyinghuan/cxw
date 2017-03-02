// Package excel 处理excel 数据
package excel

import "github.com/tealeg/xlsx"

//GeneralSheetHeader 生成表头
func GeneralSheetHeader(sheet *xlsx.Sheet, sheetHead []string) {
	row := sheet.AddRow()
	for _, head := range sheetHead {
		cell := row.AddCell()
		cell.SetString(head)
	}
}
