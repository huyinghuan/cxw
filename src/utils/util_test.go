package utils

import (
	"config"
	"fmt"
	"reflect"
	"testing"
)

func TestFieldByName(t *testing.T) {
	var s = &config.Excel{Out: "abc"}
	fmt.Println(reflect.ValueOf(s).Elem().FieldByName("Out").Interface())
}
