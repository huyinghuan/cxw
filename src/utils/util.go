package utils

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"time"
)

// FormatExcelDateToUnix 格式化excel的时间戳到 Unix
func FormatExcelDateToUnix(excelTimestamp string) (timestamp int64, err error) {
	timestamp, err = strconv.ParseInt(excelTimestamp, 10, 64)
	if err != nil {
		return
	}
	timestamp = (timestamp - 25569) * 86400
	return
}

//FormatUnixToString  Unix -> string
func FormatUnixToString(unix int64) string {
	date := time.Unix(unix, 0)
	return date.Format("2006-01-02")
}

// FormatExcelDateToString 格式化excel的时间戳到String
func FormatExcelDateToString(excelTimestamp string) (timestamp string, err error) {
	unix, err := FormatExcelDateToUnix(excelTimestamp)
	date := time.Unix(unix, 0)
	timestamp = date.Format("2006-01-02")
	return
}

//FormatStringToInt str -> 小数点4舍5入
func FormatStringToInt(value string) int {
	f, _ := strconv.ParseFloat(value, 64)
	var z = fmt.Sprintf("%.0f", f)
	j, _ := strconv.ParseInt(z, 10, 64)
	return int(j)
}

func FormatFloatToInt(f float64) int {
	var z = fmt.Sprintf("%.0f", f)
	j, _ := strconv.ParseInt(z, 10, 64)
	return int(j)
}

//FormatStringToFloat str -> float
func FormatStringToFloat(value string, precision int) (float64, error) {
	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, err
	}
	output := math.Pow(10, float64(precision))
	return float64(int(f*output+math.Copysign(0.5, f*output))) / output, nil
}

//AnythingToString anything to string
func AnythingToString(t interface{}) string {
	return fmt.Sprintf("%v", t)
}

//Invoke  struct 反射
func Invoke(any interface{}, name string, args ...interface{}) {
	inputs := make([]reflect.Value, len(args))
	for i := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	reflect.ValueOf(any).MethodByName(name).Call(inputs)
}

func FieldByName(any interface{}, name string) interface{} {
	return reflect.ValueOf(any).Elem().FieldByName(name).Interface()
}
