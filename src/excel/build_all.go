package excel

import (
	"config"
	"data"
	"fmt"
	"kind"
	"strings"

	"utils"

	"github.com/tealeg/xlsx"
)

//填充表格数据
func fillCellValue(row *xlsx.Row, cols []string, rowData *data.ExcelRowData) {
	for _, colName := range cols {
		funcName := strings.Join([]string{"CellFill", colName}, "")
		utils.Invoke(rowData, funcName, row)
	}
}

//输出一个sheet
func outputOneSheet(excelFile *xlsx.File, sheetName string, excelConfig *config.ExcelOutput, rowsData []*data.ExcelRowData) {
	var cols = excelConfig.Cols
	sheet, err := excelFile.AddSheet(sheetName)
	if err != nil {
		fmt.Printf("生成%s错误:%v \n", sheetName, err.Error())
		return
	}
	//生成表头
	GeneralSheetHeader(sheet, config.GetMapValue(cols))

	for _, rowData := range rowsData {
		//隔一行
		if rowData == nil {
			sheet.AddRow()
			continue
		}
		//添加一个行
		row := sheet.AddRow()
		//填充数据
		fillCellValue(row, cols, rowData)
	}

	GeneralSheetSum(sheet, sheetName, excelConfig, rowsData)
}

func outputSumSheet(excelConfig *config.ExcelOutput, file *xlsx.File) {
	if excelConfig.Sum == nil || len(excelConfig.Sum) == 0 {
		return
	}
	sumSheet, _ := file.AddSheet("汇总")
	allSheetsLen := len(file.Sheets)
	// -1 sumSheet -1 index=0

	//生成表头
	GeneralSheetHeader(sumSheet, config.GetMapValue(excelConfig.Sum))

	for i := 0; i < allSheetsLen-1-1; i++ {
		sheet := file.Sheets[i]
		if len(sheet.Rows) == 0 {
			continue
		}
		row := sumSheet.AddRow()
		for index, cel := range sheet.Rows[len(sheet.Rows)-1].Cells {

			if cel.Value == "" && excelConfig.Sum[index] == "" {
				continue
			}

			if cel.Value == "" && excelConfig.Sum[index] != "" {
				newCel := row.AddCell()
				newCel.SetValue(sheet.Rows[len(sheet.Rows)-2].Cells[index].Value)
			}

			if cel.Value != "" {
				newCel := row.AddCell()
				*newCel = *cel
			}
		}
	}

}

func outputOneExcel(configFileName string, outputFile string, excelRows []*data.ExcelRowData) {
	//为了能够反射struct的函数，new个空间
	var excelFactory = new(kind.Factory)
	//获取表格配置文件
	var excelConfig = config.ReadExcelOutput("config/excel/" + configFileName + ".json")

	//------------------------------------
	//获取 配置里面的 数据分组函数名
	//通过反射机制 加工分组 数据
	categoryGroup := make(map[string][]*data.ExcelRowData)
	utils.Invoke(excelFactory, excelConfig.ExcelKind, excelRows, excelConfig, categoryGroup)
	//强转
	//------------------------------------

	file := xlsx.NewFile()
	//生产sheet
	for sheetName, rowsData := range categoryGroup {
		outputOneSheet(file, sheetName, excelConfig, rowsData)
	}

	outputSumSheet(excelConfig, file)

	//存文件
	err := file.Save(outputFile)
	if err != nil {
		fmt.Printf("\n错误信息:\n%v\n\n", err.Error())
		fmt.Printf("表格 %s 生成失败\n", outputFile)
	} else {
		fmt.Printf("表格 %s 生成成功\n", outputFile)
	}

}

//BuildOne 输入一个数据源，输出一个表格， 包含多个sheet
func BuildOne(inputDataSource string, sheetConfig map[string]string, layout *config.ExcelLayout) {
	//获取数据源的数据
	var excelRowsData = GetDataSource(inputDataSource, layout)
	if err := kind.CalculateAll(excelRowsData, layout.Skip); err != nil {
		return
	}
	for configFileName, outputFile := range sheetConfig {
		outputOneExcel(configFileName, outputFile, excelRowsData)
	}
}

//BuildAll 生成所有文件
func BuildAll(outputConfig map[string]map[string]string, layout *config.ExcelLayout) {
	for key, value := range outputConfig {
		BuildOne(key, value, layout)
	}
}
