package kind

import (
	"data"
	"fmt"
	"utils"
)

func vertifyMileageAndOilPriceAndOilWear(lineData *data.ExcelRowData, index int) int {
	if lineData.Mileage != 0 && (lineData.OilPrice == 0 || lineData.OilWear == 0) {
		fmt.Printf("源数据校验 %s 错误,错误行数%d\n", "==油价或油耗==", index)
		return 0
	}
	return 1
}

func vertifyOilRefuelNumberAndOilStationPrice(lineData *data.ExcelRowData, index int) int {
	if (lineData.OilRefuelNumber == 0 && lineData.OilStationPrice != 0) ||
		(lineData.OilRefuelNumber != 0 && lineData.OilStationPrice == 0) {
		fmt.Printf("源数据校验 %s 错误,行数: %d\n", "==加油站加油数量或加油站油价==", index)
		return 0
	}
	return 1
}

func vertify(lineData *data.ExcelRowData, index int) int {
	return vertifyMileageAndOilPriceAndOilWear(lineData, index) * vertifyOilRefuelNumberAndOilStationPrice(lineData, index)
}

//CalculateAll 计算应该计算的东西
func CalculateAll(lineDataArrar []*data.ExcelRowData, lineSkipNumber int) error {
	var vertifyFlag = true
	for index, lineData := range lineDataArrar {
		if vertify(lineData, index+lineSkipNumber+1) != 1 {
			vertifyFlag = false
			continue
		}
		//自己加油
		OilSelfCosts := (float64(lineData.Mileage) * lineData.OilWear) / 100 * lineData.OilPrice
		//油站加油
		lineData.OilStationCosts = utils.FormatFloatToInt(lineData.OilRefuelNumber * lineData.OilStationPrice)
		lineData.ReturnMoney = utils.FormatFloatToInt(OilSelfCosts - float64(lineData.OilStationCosts))
		lineData.OilSelfCosts = utils.FormatFloatToInt(OilSelfCosts)

		//数据单行计算
		lineData.ResultForOilReturnCash = float64(lineData.ReturnMoney) + lineData.TransitHighwayCashCosts + float64(lineData.RepairCashCosts)
	}
	if !vertifyFlag {
		return fmt.Errorf("数据校验错误！")
	}
	return nil
}
