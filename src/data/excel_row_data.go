package data

// ExcelRowData 线路数据
type ExcelRowData struct {
	SubmitDate              int64   //报账日期
	DepartDate              int64   //发车日期
	CarLine                 string  //路线
	CarLineType             string  //路线类型
	Payer                   string  //结算方
	CarID                   string  //车牌
	CarType                 string  //车型
	Mileage                 int     //里程
	Times                   int     //趟数
	OilWear                 float64 //油耗
	OilPrice                float64 //油价
	OilSelfCosts            int     //自己加油的费用 （Mileage * OilWear)/100 * OilPrice
	OilRefuelNumber         float64 //加油量/公升
	OilStationPrice         float64 //加油站油价
	OilStationCosts         int     //加油站加油的费用 (OilRefuelNumber * OilStationPrice)
	ReturnMoney             int     //返现 OilSelfCosts - OilStationCosts
	TransitHighwayCashCosts float64 //现金过境费
	TransitHighwayETCCosts  int     //ETC过境费
	RepairCashCosts         int     //实报实修修理费
	RepairDepreciationRate  float64 //修理费返现比例
	RepairDepreciationCosts int     //修理费返现  (RepairDepreciationRate * Mileage)
	Maintain                int     //保养费用
	Material                int     //材料费用
	Fine                    int     //罚款
	FreightCode             string  //货运编号 封铅
	CarAndGoodsWeight       int     //总重
	CarOnlyWeight           int     //皮重
	GoodsWeight             int     //净重
	RentCosts               int     //租外面的车
	Other                   int     //其他
	Remark                  string  //备注
	Total                   int     //总计
	//------------   以下为计算结果
	ResultForOilReturnCash float64 //油费返现
}
