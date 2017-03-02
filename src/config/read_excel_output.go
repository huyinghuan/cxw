package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//ReadExcelOutput 读取一个输出excel的文件格式
func ReadExcelOutput(outputConfigName string) (outputConfig *ExcelOutput) {

	fileContentByte, err := ioutil.ReadFile(outputConfigName)
	if err != nil {
		fmt.Printf("%v\n读输出excel %s 配置出错。", err, outputConfigName)
		panic("Error!")
	}

	if err := json.Unmarshal(fileContentByte, &outputConfig); err != nil {
		fmt.Printf("%v\n解析配置出错:%s", err, outputConfigName)
		panic("Error!")
	}
	return outputConfig
}
