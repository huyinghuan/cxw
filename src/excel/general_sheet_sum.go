// Package excel 处理excel 数据
package excel

import (
	"config"
	"data"
	"reflect"
	"strings"
	"utils"

	"fmt"

	"github.com/tealeg/xlsx"
)

//GeneralSheetSum 生成sheet合计
func GeneralSheetSum(sheet *xlsx.Sheet, sheetName string, excelConfig *config.ExcelOutput, rowsData []*data.ExcelRowData) {
	sheetSum := excelConfig.Sum
	if sheetSum == nil || len(sheetSum) == 0 {
		return
	}
	sumRow := sheet.AddRow()
	for _, sumField := range sheetSum {
		if sumField == "" || strings.Contains(sumField, "~") {
			sumRow.AddCell().SetValue("")
			continue
		}
		sum := float64(0)
		for index, rowData := range rowsData {
			valueReflect := utils.FieldByName(rowData, sumField)

			switch reflect.TypeOf(valueReflect).String() {
			case "int":
				sum = sum + float64(valueReflect.(int))
			case "float64":
				sum = sum + valueReflect.(float64)
			default:
				fmt.Printf("Sheet %s, 第 %d 行, %s 不能正确识别数据类型", sheetName, index+1, sumField)
			}
		}
		sumRow.AddCell().SetFloat(sum)
	}
}
