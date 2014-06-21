package helper

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

//返回現在時間 mysql datetime
//格式:2014-12-21 21:20:32
func NowTime() string {
	t := time.Now()
	y, m, d := t.Date()
	H := t.Hour()
	i := t.Minute()
	se := t.Second()
	return fmt.Sprintf("%4d-%02d-%02d %02d:%02d:%02d", y, m, d, H, i, se)
}

// 轉換成int64
func ToInt(i interface{}) int {
	kind := reflect.TypeOf(i).Kind()
	if kind == reflect.String {
		v, err := strconv.Atoi(i.(string))
		if err == nil {
			return int(v)
		}
	}
	if kind == reflect.Int {
		return int(i.(int))
	}
	if kind == reflect.Int8 {
		return int(i.(int8))
	}
	if kind == reflect.Int16 {
		return int(i.(int16))
	}
	if kind == reflect.Int32 {
		return int(i.(int32))
	}
	if kind == reflect.Int64 {
		return int(i.(int64))
	}
	if kind == reflect.Uint {
		return int(i.(uint))
	}
	if kind == reflect.Uint8 {
		return int(i.(uint8))
	}
	if kind == reflect.Uint16 {
		return int(i.(uint16))
	}
	if kind == reflect.Uint32 {
		return int(i.(uint32))
	}
	if kind == reflect.Uint64 {
		return int(i.(uint64))
	}
	if kind == reflect.Float32 {
		return int(i.(float32))
	}
	if kind == reflect.Float64 {
		return int(i.(float64))
	}
	if kind == reflect.Bool {
		if i.(bool) == true {
			return int(1)
		} else {
			return int(0)
		}
	}
	return int(0)
}

func ToString(s interface{}) string {
	kind := reflect.TypeOf(s).Kind()
	if kind == reflect.String {
		return s.(string)
	}
	if kind == reflect.Int {
		return strconv.Itoa(s.(int))
	}
	if kind == reflect.Int8 {
		return strconv.Itoa(int(s.(int8)))
	}
	if kind == reflect.Int16 {
		return strconv.Itoa(int(s.(int16)))
	}
	if kind == reflect.Int32 {
		return strconv.Itoa(int(s.(int32)))
	}
	if kind == reflect.Int64 {
		return strconv.Itoa(int(s.(int64)))
	}
	if kind == reflect.Uint {
		return strconv.Itoa(int(s.(uint)))
	}
	if kind == reflect.Uint8 {
		return strconv.Itoa(int(s.(uint8)))
	}
	if kind == reflect.Uint16 {
		return strconv.Itoa(int(s.(uint16)))
	}
	if kind == reflect.Uint32 {
		return strconv.Itoa(int(s.(uint32)))
	}
	if kind == reflect.Uint64 {
		return strconv.Itoa(int(s.(uint64)))
	}
	return ""
}
