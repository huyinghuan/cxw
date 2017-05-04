package kind

import (
	"config"
	"data"
	"utils"
)

func doCategory(excelRows []*data.ExcelRowData, category string, lineData map[string][]*data.ExcelRowData) {
	for _, item := range excelRows {
		categoryReflectValue := utils.FieldByName(item, category)
		categoryValue := categoryReflectValue.(string)
		if _, ok := lineData[categoryValue]; !ok {
			lineData[categoryValue] = []*data.ExcelRowData{}
		}
		lineData[categoryValue] = append(lineData[categoryValue], item)
	}
}

//mapToArray 分割map成array，中间隔几行
func mapToArray(lineData map[string][]*data.ExcelRowData, split int) []*data.ExcelRowData {
	sheetRows := []*data.ExcelRowData{}
	for _, value := range lineData {
		sheetRows = append(sheetRows, value...)
	}
	return sheetRows
}

//SimpleCategory 简单的数据数据分类
//excelRows 数据
//category 分类依据
func (excelKind *Factory) SimpleCategory(excelRows []*data.ExcelRowData, excelOutputConfig *config.ExcelOutput, lineData map[string][]*data.ExcelRowData) {
	var category = excelOutputConfig.Category
	doCategory(excelRows, category, lineData)
}

//TwolayerCategory 简单的数据数据两层分类
//excelRows 数据
//category 分类依据
func (excelKind *Factory) TwolayerCategory(excelRows []*data.ExcelRowData, excelOutputConfig *config.ExcelOutput, lineData map[string][]*data.ExcelRowData) {
	var category = excelOutputConfig.Category
	inExcelCategory := make(map[string][]*data.ExcelRowData)
	doCategory(excelRows, category, inExcelCategory)
	//已经分类到数据再次分类
	// _ sheetName
	// categoryData sheetRows
	for sheetName, categoryData := range inExcelCategory {
		//sheet内进行分类
		inSheetCategory := make(map[string][]*data.ExcelRowData)
		doCategory(categoryData, excelOutputConfig.TwoCategory, inSheetCategory)
		lineData[sheetName] = mapToArray(inSheetCategory, 1)
	}

}

func (excelKind *Factory) MixCategory(excelRows []*data.ExcelRowData, excelOutputConfig *config.ExcelOutput, lineData map[string][]*data.ExcelRowData) {
	var category = excelOutputConfig.Category
	var mixto = excelOutputConfig.Mixto

	for _, item := range excelRows {
		categoryReflectValue := utils.FieldByName(item, category)
		mixtoRelectValue := utils.FieldByName(item, mixto)
		categoryValue := categoryReflectValue.(string)
		mixtoValue := mixtoRelectValue.(string)

		//hasher := md5.New()
		//hasher.Write([]byte(categoryValue + mixtoValue))
		key := categoryValue + mixtoValue
		if _, ok := lineData[key]; !ok {
			lineData[key] = []*data.ExcelRowData{}
		}
		lineData[key] = append(lineData[key], item)
	}
}
