package config

import (
	"fmt"
	"testing"
)

func z(zi map[string]map[string]string) {
	for key := range zi {
		fmt.Printf("key: %s \n", key)
	}
}

func TestReadOutputYaml(t *testing.T) {
	if i, err := ReadOutputYaml("/Users/hyh/mywork/CXW/config/output.yml"); err != nil {
		t.Fail()
	} else {
		z(i)
	}

}
