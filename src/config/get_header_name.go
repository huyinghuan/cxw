package config

import (
	"fmt"
	"io/ioutil"
	"strings"

	_yaml "gopkg.in/yaml.v2"
)

//GetMapValue 获取列名数组
func GetMapValue(keys []string) []string {
	fileContentByte, err := ioutil.ReadFile("config/name.yml")
	if err != nil {
		fmt.Println(err)
		panic("read name.json error")
	}
	var cols = make(map[string]string)
	err = _yaml.Unmarshal(fileContentByte, &cols)
	if err != nil {
		fmt.Println(err)
		panic("json parse name.json error")
	}
	values := []string{}
	for _, key := range keys {
		//兼容sum字段
		key = strings.Replace(key, "~", "", -1)
		if key == "" {
			continue
		}
		value := cols[key]
		if value == "" {
			value = key
		}
		values = append(values, value)
	}
	return values
}
