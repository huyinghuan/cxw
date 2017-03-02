package config

import (
	"fmt"
	"io/ioutil"

	_yaml "gopkg.in/yaml.v2"
)

//ReadOutputYaml 读取配置文件 返回map类型
func ReadOutputYaml(filename string) (jsonMap map[string]map[string]string, readErr error) {
	fileContentByte, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("%v\n读取输出总配置出错。\n", err)
		readErr = err
		return
	}
	if err := _yaml.Unmarshal(fileContentByte, &jsonMap); err != nil {
		fmt.Printf("%v\n解析配置出错:%s\n", err, filename)
		readErr = err
		return
	}
	return
}
