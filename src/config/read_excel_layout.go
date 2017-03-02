package config

import (
	"fmt"
	"io/ioutil"

	_yaml "gopkg.in/yaml.v2"
)

//ReadExcelLayout 读取数据源文件的 excel 结构
func ReadExcelLayout(filename string) (layout *ExcelLayout, returnErr error) {
	fileContentByte, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		returnErr = err
		return
	}
	err = _yaml.Unmarshal(fileContentByte, &layout)
	return
}
