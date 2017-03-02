package excel

import (
	"config"
	"data"
	"fmt"
	"os"
	"strings"
	"utils"

	_xlsx "github.com/tealeg/xlsx"
)

//analyzRow 分析行数据
func analyzRow(row *_xlsx.Row, cellOrder []string) (rowData *data.ExcelRowData, err error) {
	var cels = row.Cells

	if len(cels) == 0 || cels[0].Value == "" {
		return
	}
	var data = new(data.ExcelRowData)
	var celsLength = len(cels)
	//fmt.Println(celsLength)
	//fmt.Println(len(cellOrder))
	// if len(cellOrder) > celsLength {
	// 	return nil, fmt.Errorf("表格列配置错误,请对照数据源excel和 excel-layout.json的配置顺序，数目")
	// }
	for i := 0; i < len(cellOrder); i++ {
		if i > celsLength-1 {
			continue
		}
		field := cellOrder[i]
		if field == "" {
			continue
		}
		funcName := strings.Join([]string{"Set", field}, "")
		utils.Invoke(data, funcName, cels[i])
	}
	return data, nil
}

// GetDataSource 根据excel文件名获取 struct slice
func GetDataSource(excelName string, excelLayout *config.ExcelLayout) []*data.ExcelRowData {
	var i = excelLayout.Skip
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			fmt.Printf("\n出错行数 %d\n", i+1)
			os.Exit(1)
		}
	}()
	//打开excel文件
	excel, err := _xlsx.OpenFile(excelName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var sheets = excel.Sheets
	var rows = sheets[0].Rows
	var dataArray []*data.ExcelRowData
	for ; i < len(rows); i++ {
		item, analyzRowErr := analyzRow(rows[i], excelLayout.Order)
		if analyzRowErr != nil {
			panic(analyzRowErr)
		}
		if item == nil {
			continue
		}
		dataArray = append(dataArray, item)
	}
	return dataArray
}
