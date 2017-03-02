package config

import (
	"fmt"
)

var VersionMap = map[string]string{
	"version":   "1.2",
	"copyright": "ec.huyinghuan@gmail.com",
	"publish":   "2017-03-02",
	"name":      "申通物流 数据分析系统【春晓版】"}

//PrintStatement 打印版权声明
func PrintStatement() {
	fmt.Printf("============== %s ============\n", VersionMap["name"])
	fmt.Printf("  版本: %s   发布时间: %s  \n", VersionMap["version"], VersionMap["publish"])
	fmt.Printf("  版权所有: %s \n", VersionMap["copyright"])
	fmt.Printf("  特别鸣谢: %s \n", "宋春晓")
	fmt.Println("===========================================================")
	fmt.Printf("\n\n")
}
