package main

import (
	"config"
	"excel"
	"fmt"
	"net/http"
)

//保持进程存在
func itHere() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {})
	http.ListenAndServe("localhost:15200", nil)
}

func main() {
	//打印声明
	config.PrintStatement()
	//获取数据源格式
	excelLayout, err := config.ReadExcelLayout("config/excel-layout.yml")
	if err != nil {
		fmt.Println("读取数据源配置文件错误")
		return
	}

	outputConfig, err := config.ReadOutputYaml("config/output.yml")
	if err != nil {
		return
	}

	fmt.Println("=====> 数据分析进行中.....")
	//生成所有文件
	excel.BuildAll(outputConfig, excelLayout)

	fmt.Println("=====> 数据分析结束")
	// itHere()
}
