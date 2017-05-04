package config

type ExcelOutput struct {
	ExcelKind   string
	Category    string
	TwoCategory string
	Cols        []string
	Sum         []string
	Mixto       string
	Onlysum     bool   //只要汇总表
	Calculate   string //计算公式
}
