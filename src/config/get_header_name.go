package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

//GetMapValue 获取列名数组
func GetMapValue(keys []string) []string {
	fileContentByte, err := ioutil.ReadFile("config/name.json")
	if err != nil {
		fmt.Println(err)
		panic("read name.json error")
	}
	var cols = make(map[string]string)
	err = json.Unmarshal(fileContentByte, &cols)
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
