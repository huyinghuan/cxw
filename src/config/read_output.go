package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//ReadOutput 读取配置文件 返回map类型
func ReadOutput(filename string) (jsonMap map[string]interface{}, readErr error) {
	fileContentByte, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("%v\n读取输出总配置出错。", err)
		readErr = err
		return
	}
	var temp interface{}
	if err := json.Unmarshal(fileContentByte, &temp); err != nil {
		fmt.Printf("%v\n解析配置出错:%s", err, filename)
		readErr = err
		return
	}
	return temp.(map[string]interface{}), nil
}
