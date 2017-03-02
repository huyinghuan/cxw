package data

import (
	"fmt"
	"strconv"
	"strings"
	"utils"

	"github.com/tealeg/xlsx"
)

//SetSubmitDate 处理报账日期
func (lineData *ExcelRowData) SetSubmitDate(cell *xlsx.Cell) {
	value := cell.Value
	timestamp, err := utils.FormatExcelDateToUnix(value)
	if err != nil {
		fmt.Printf("格式化报账日期错误：%v", err)
		panic("parse error")
	} else {
		lineData.SubmitDate = timestamp
	}

}

//SetDepartDate 处理发车日期
func (lineData *ExcelRowData) SetDepartDate(cell *xlsx.Cell) {
	value := cell.Value
	timestamp, err := utils.FormatExcelDateToUnix(value)
	if err != nil {
		fmt.Printf("格式化报账日期错误：%v", err)
		panic("parse error")
	} else {
		lineData.DepartDate = timestamp
	}
}

//SetCarLine 处理路线
func (lineData *ExcelRowData) SetCarLine(cell *xlsx.Cell) {
	value := cell.Value
	lineData.CarLine = strings.TrimSpace(value)
}

//SetCarLineType 路线类型
func (lineData *ExcelRowData) SetCarLineType(cell *xlsx.Cell) {
	value := cell.Value
	lineData.CarLineType = strings.TrimSpace(value)
}

//SetPayer 结算方
func (lineData *ExcelRowData) SetPayer(cell *xlsx.Cell) {
	value := cell.Value
	lineData.Payer = strings.TrimSpace(value)
}

//SetCarID 处理车牌
func (lineData *ExcelRowData) SetCarID(cell *xlsx.Cell) {
	value := cell.Value
	lineData.CarID = strings.TrimSpace(value)
}

//SetCarType 处理车型
func (lineData *ExcelRowData) SetCarType(cell *xlsx.Cell) {
	value := cell.Value
	lineData.CarType = strings.TrimSpace(value)
}

//SetMileage 处理里程
func (lineData *ExcelRowData) SetMileage(cell *xlsx.Cell) {
	value := cell.Value
	if value == "" {
		lineData.Mileage = 0
		return
	}
	mileage, err := strconv.Atoi(value)
	if err != nil {
		fmt.Printf("格式化%s错误：%v", "里程", err)
		panic("parse error")
	} else {
		lineData.Mileage = mileage
	}
}

//SetTimes 有效趟数
func (lineData *ExcelRowData) SetTimes(cell *xlsx.Cell) {
	value := cell.Value
	if value == "" {
		lineData.Times = 0
		return
	}
	times, err := strconv.Atoi(value)
	if err != nil {
		fmt.Printf("格式化趟数错误：%v", err)
		panic("parse error")
	} else {
		lineData.Times = times
	}
}

//SetOilWear 处理 百公里油耗
func (lineData *ExcelRowData) SetOilWear(cell *xlsx.Cell) {
	value := cell.Value
	if value == "" {
		lineData.OilWear = 0
		return
	}
	oilWear, err := strconv.ParseFloat(value, 64)
	if err != nil {
		fmt.Printf("格式化%s错误：%v", "百公里油耗", err)
		panic("parse error")
	} else {
		lineData.OilWear = oilWear
	}
}

//SetOilPrice 油价
func (lineData *ExcelRowData) SetOilPrice(cell *xlsx.Cell) {
	value := cell.Value
	if value == "" {
		lineData.OilPrice = 0
		return
	}
	OilPrice, err := strconv.ParseFloat(value, 64)
	if err != nil {
		fmt.Printf("格式化油价错误：%v", err)
		panic("parse error")
	} else {
		lineData.OilPrice = OilPrice
	}
}

//SetOilRefuelNumber 加油的公升
func (lineData *ExcelRowData) SetOilRefuelNumber(cell *xlsx.Cell) {
	value := cell.Value
	if value == "" {
		lineData.OilRefuelNumber = 0
		return
	}
	oilRefuelNumber, err := utils.FormatStringToFloat(value, 2)
	if err != nil {
		fmt.Printf("格式化%s错误：%v", "加油的公升", err)
		panic("parse error")
	} else {
		lineData.OilRefuelNumber = oilRefuelNumber
	}
}

//SetOilStationPrice 加油站加油单价
func (lineData *ExcelRowData) SetOilStationPrice(cell *xlsx.Cell) {
	value := cell.Value
	if value == "" {
		lineData.OilStationPrice = 0
		return
	}
	OilStationPrice, err := utils.FormatStringToFloat(value, 2)
	if err != nil {
		fmt.Printf("格式化加油钱错误：%v", err)
		panic("parse error")
	} else {
		lineData.OilStationPrice = OilStationPrice
	}
}

//SetReturnMoney  返现
func (lineData *ExcelRowData) SetReturnMoney(cell *xlsx.Cell) {
	value := cell.Value
	if value == "" {
		lineData.ReturnMoney = 0
		return
	}
	returnMoney, err := strconv.Atoi(value)
	if err != nil {
		fmt.Printf("格式化%s错误：%v", "返现", err)
		panic("parse error")
	} else {
		lineData.ReturnMoney = returnMoney
	}
}

//SetTransitHighwayCashCosts 现金过境费
func (lineData *ExcelRowData) SetTransitHighwayCashCosts(cell *xlsx.Cell) {
	value := cell.Value
	if value == "" {
		lineData.TransitHighwayCashCosts = 0
		return
	}
	TransitHighwayCashCosts, err := utils.FormatStringToFloat(value, 2)
	if err != nil {
		fmt.Printf("格式化现金过境费错误：%v", err)
		panic("parse error")
	} else {
		lineData.TransitHighwayCashCosts = TransitHighwayCashCosts
	}
}

//SetTransitHighwayETCCosts ETC过境费
func (lineData *ExcelRowData) SetTransitHighwayETCCosts(cell *xlsx.Cell) {
	value := cell.Value
	if value == "" {
		lineData.TransitHighwayETCCosts = 0
		return
	}
	TransitHighwayETCCosts, err := strconv.Atoi(value)
	if err != nil {
		fmt.Printf("格式化%s错误：%v", "ETC过境费", err)
		panic("parse error")
	} else {
		lineData.TransitHighwayETCCosts = TransitHighwayETCCosts
	}
}

//SetRepairCashCosts 修理费
func (lineData *ExcelRowData) SetRepairCashCosts(cell *xlsx.Cell) {
	value := cell.Value
	if value == "" {
		lineData.RepairCashCosts = 0
		return
	}
	RepairCashCosts, err := strconv.Atoi(value)
	if err != nil {
		fmt.Printf("格式化%s错误：%v", "修理费", err)
		panic("parse error")
	} else {
		lineData.RepairCashCosts = RepairCashCosts
	}
}

// SetRepairDepreciationRate 修理费返现
func (lineData *ExcelRowData) SetRepairDepreciationRate(cell *xlsx.Cell) {
	value := cell.Value
	if value == "" {
		lineData.RepairDepreciationRate = 0
		return
	}
	RepairDepreciationRate, err := utils.FormatStringToFloat(value, 4)
	if err != nil {
		fmt.Printf("格式化%s错误：%v", "修理费折旧比例", err)
		panic("parse error")
	} else {
		lineData.RepairDepreciationRate = RepairDepreciationRate
	}
}

//SetFreightCode 封铅
func (lineData *ExcelRowData) SetFreightCode(cell *xlsx.Cell) {
	value := cell.Value
	lineData.FreightCode = strings.TrimSpace(value)
}

//SetCarAndGoodsWeight 总重
func (lineData *ExcelRowData) SetCarAndGoodsWeight(cell *xlsx.Cell) {
	value := cell.Value
	if value == "" {
		lineData.CarAndGoodsWeight = 0
		return
	}
	CarAndGoodsWeight, err := strconv.Atoi(value)
	if err != nil {
		fmt.Printf("格式化%s错误：%v", "总重", err)
		panic("parse error")
	} else {
		lineData.CarAndGoodsWeight = CarAndGoodsWeight
	}
}

//SetCarOnlyWeight 皮重
func (lineData *ExcelRowData) SetCarOnlyWeight(cell *xlsx.Cell) {
	value := cell.Value
	if value == "" {
		lineData.CarOnlyWeight = 0
		return
	}
	CarOnlyWeight, err := strconv.Atoi(value)
	if err != nil {
		fmt.Printf("格式化%s错误：%v", "皮重", err)
		panic("parse error")
	} else {
		lineData.CarOnlyWeight = CarOnlyWeight
	}
}

//SetGoodsWeight 净重
func (lineData *ExcelRowData) SetGoodsWeight(cell *xlsx.Cell) {
	value := cell.Value
	if value == "" {
		lineData.GoodsWeight = 0
		return
	}
	GoodsWeight, err := strconv.Atoi(value)
	if err != nil {
		fmt.Printf("格式化%s错误：%v", "净重", err)
		panic("parse error")
	} else {
		lineData.GoodsWeight = GoodsWeight
	}
}

//SetRemark 备注
func (lineData *ExcelRowData) SetRemark(cell *xlsx.Cell) {
	value := cell.Value
	lineData.Remark = strings.TrimSpace(value)
}

//SetMaintain 保养
func (lineData *ExcelRowData) SetMaintain(cell *xlsx.Cell) {
	value := cell.Value
	if value == "" {
		lineData.Maintain = 0
		return
	}
	Maintain, err := strconv.Atoi(value)
	if err != nil {
		fmt.Printf("格式化%s错误：%v", "保养", err)
		panic("parse error")
	} else {
		lineData.Maintain = Maintain
	}
}

//SetMaterial 材料
func (lineData *ExcelRowData) SetMaterial(cell *xlsx.Cell) {
	value := cell.Value
	if value == "" {
		lineData.Material = 0
		return
	}
	Material, err := strconv.Atoi(value)
	if err != nil {
		fmt.Printf("格式化%s错误：%v", "材料", err)
		panic("parse error")
	} else {
		lineData.Material = Material
	}
}

//SetFine 罚款
func (lineData *ExcelRowData) SetFine(cell *xlsx.Cell) {
	value := cell.Value
	if value == "" {
		lineData.Fine = 0
		return
	}
	Fine, err := strconv.Atoi(value)
	if err != nil {
		fmt.Printf("格式化%s错误：%v", "罚款", err)
		panic("parse error")
	} else {
		lineData.Fine = Fine
	}
}

//SetRentCosts 租用车
func (lineData *ExcelRowData) SetRentCosts(cell *xlsx.Cell) {
	value := cell.Value
	if value == "" {
		lineData.RentCosts = 0
		return
	}
	RentCosts, err := strconv.Atoi(value)
	if err != nil {
		fmt.Printf("格式化%s错误：%v", "租用车", err)
		panic("parse error")
	} else {
		lineData.RentCosts = RentCosts
	}
}

//SetOther 租用车
func (lineData *ExcelRowData) SetOther(cell *xlsx.Cell) {
	value := cell.Value
	if value == "" {
		lineData.Other = 0
		return
	}
	Other, err := strconv.Atoi(value)
	if err != nil {
		fmt.Printf("格式化%s错误：%v", "租用车", err)
		panic("parse error")
	} else {
		lineData.Other = Other
	}
}
