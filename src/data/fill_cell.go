package data

import (
	"utils"

	"github.com/tealeg/xlsx"
)

//CellFillSubmitDate 填充报账日期
func (rowData *ExcelRowData) CellFillSubmitDate(row *xlsx.Row) {
	cell := row.AddCell()
	//cell.SetDateTime(time.Unix(rowData.SubmitDate, 0))
	cell.SetString(utils.FormatUnixToString(rowData.SubmitDate))
}

func (rowData *ExcelRowData) CellFillDepartDate(row *xlsx.Row) {
	cell := row.AddCell()
	//cell.SetDateTime(time.Unix(rowData.DepartDate, 0))
	cell.SetString(utils.FormatUnixToString(rowData.DepartDate))
}

func (rowData *ExcelRowData) CellFillCarID(row *xlsx.Row) {
	cell := row.AddCell()
	cell.SetString(rowData.CarID)
}

func (rowData *ExcelRowData) CellFillCarType(row *xlsx.Row) {
	cell := row.AddCell()
	cell.SetString(rowData.CarType)
}

func (rowData *ExcelRowData) CellFillCarLine(row *xlsx.Row) {
	cell := row.AddCell()
	cell.SetString(rowData.CarLine)
}
func (rowData *ExcelRowData) CellFillCarLineType(row *xlsx.Row) {
	cell := row.AddCell()
	cell.SetString(rowData.CarLineType)
}
func (rowData *ExcelRowData) CellFillPayer(row *xlsx.Row) {
	cell := row.AddCell()
	cell.SetString(rowData.Payer)
}
func (rowData *ExcelRowData) CellFillMileage(row *xlsx.Row) {
	cell := row.AddCell()
	cell.SetInt(rowData.Mileage)
}
func (rowData *ExcelRowData) CellFillTimes(row *xlsx.Row) {
	cell := row.AddCell()
	cell.SetInt(rowData.Times)
}
func (rowData *ExcelRowData) CellFillOilWear(row *xlsx.Row) {
	cell := row.AddCell()
	cell.SetFloat(rowData.OilWear)
}
func (rowData *ExcelRowData) CellFillOilPrice(row *xlsx.Row) {
	cell := row.AddCell()
	cell.SetFloat(rowData.OilPrice)
}
func (rowData *ExcelRowData) CellFillOilSelfCosts(row *xlsx.Row) {
	cell := row.AddCell()
	cell.SetInt(rowData.OilSelfCosts)
}
func (rowData *ExcelRowData) CellFillOilRefuelNumber(row *xlsx.Row) {
	cell := row.AddCell()
	cell.SetFloat(rowData.OilRefuelNumber)
}
func (rowData *ExcelRowData) CellFillOilStationPrice(row *xlsx.Row) {
	cell := row.AddCell()
	cell.SetFloat(rowData.OilStationPrice)
}
func (rowData *ExcelRowData) CellFillOilStationCosts(row *xlsx.Row) {
	cell := row.AddCell()
	cell.SetInt(rowData.OilStationCosts)
}
func (rowData *ExcelRowData) CellFillReturnMoney(row *xlsx.Row) {
	cell := row.AddCell()
	cell.SetInt(rowData.ReturnMoney)
}
func (rowData *ExcelRowData) CellFillTransitHighwayCashCosts(row *xlsx.Row) {
	cell := row.AddCell()
	cell.SetFloat(rowData.TransitHighwayCashCosts)
}
func (rowData *ExcelRowData) CellFillTransitHighwayETCCosts(row *xlsx.Row) {
	cell := row.AddCell()
	cell.SetInt(rowData.TransitHighwayETCCosts)
}
func (rowData *ExcelRowData) CellFillRepairCashCosts(row *xlsx.Row) {
	cell := row.AddCell()
	cell.SetInt(rowData.RepairCashCosts)
}

func (rowData *ExcelRowData) CellFillRepairDepreciationRate(row *xlsx.Row) {
	cell := row.AddCell()
	cell.SetFloat(rowData.RepairDepreciationRate)
}
func (rowData *ExcelRowData) CellFillMaintain(row *xlsx.Row) {
	cell := row.AddCell()
	cell.SetInt(rowData.Maintain)
}
func (rowData *ExcelRowData) CellFillMaterial(row *xlsx.Row) {
	cell := row.AddCell()
	cell.SetInt(rowData.Material)
}

func (rowData *ExcelRowData) CellFillFine(row *xlsx.Row) {
	cell := row.AddCell()
	cell.SetInt(rowData.Fine)
}
func (rowData *ExcelRowData) CellFillFreightCode(row *xlsx.Row) {
	cell := row.AddCell()
	cell.SetString(rowData.FreightCode)
}
func (rowData *ExcelRowData) CellFillCarAndGoodsWeight(row *xlsx.Row) {
	cell := row.AddCell()
	cell.SetInt(rowData.CarAndGoodsWeight)
}
func (rowData *ExcelRowData) CellFillCarOnlyWeight(row *xlsx.Row) {
	cell := row.AddCell()
	cell.SetInt(rowData.CarOnlyWeight)
}
func (rowData *ExcelRowData) CellFillGoodsWeight(row *xlsx.Row) {
	cell := row.AddCell()
	cell.SetInt(rowData.GoodsWeight)
}
func (rowData *ExcelRowData) CellFillRentCosts(row *xlsx.Row) {
	cell := row.AddCell()
	cell.SetInt(rowData.RentCosts)
}
func (rowData *ExcelRowData) CellFillOther(row *xlsx.Row) {
	cell := row.AddCell()
	cell.SetInt(rowData.Other)
}
func (rowData *ExcelRowData) CellFillRemark(row *xlsx.Row) {
	cell := row.AddCell()
	cell.SetString(rowData.Remark)
}
func (rowData *ExcelRowData) CellFillTotal(row *xlsx.Row) {
	cell := row.AddCell()
	cell.SetInt(rowData.Total)
}

//CellFillResultForOilReturnCash 邮费返现
func (rowData *ExcelRowData) CellFillResultForOilReturnCash(row *xlsx.Row) {
	cell := row.AddCell()
	cell.SetFloat(rowData.ResultForOilReturnCash)
}
