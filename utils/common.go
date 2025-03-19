package utils

import (
	"encoding/json"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func IntValue(str string, defaultValue int) int {
	value, err := strconv.Atoi(str)
	if err != nil {
		return defaultValue
	}
	return value
}

func InterfaceToString(value interface{}) string {
	if str, ok := value.(string); ok {
		return str
	}
	return JsonEncode(value)
}
func JsonEncode(value interface{}) string {
	bytes, err := json.Marshal(value)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func JsonDecode(src string, dest interface{}) error {
	return json.Unmarshal([]byte(src), dest)
}

func FloatValue(str string) float64 {
	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}
	return value
}

func BoolValue(str string) bool {
	value, err := strconv.ParseBool(str)
	if err != nil {
		return false
	}
	return value
}

// CopyObject 拷贝对象
func CopyObject(src interface{}, dst interface{}) error {

	srcType := reflect.TypeOf(src)
	srcValue := reflect.ValueOf(src)
	dstValue := reflect.ValueOf(dst).Elem()
	reflect.TypeOf(dst)
	for i := 0; i < srcType.NumField(); i++ {
		field := srcType.Field(i)
		value := dstValue.FieldByName(field.Name)
		if !value.IsValid() {
			continue
		}
		// 数据类型相同，直接赋值
		v := srcValue.FieldByName(field.Name)
		if value.Type() == field.Type {
			value.Set(v)
		} else {
			// src data type is  string，dst data type is slice, map, struct
			// use json decode the data
			if field.Type.Kind() == reflect.String && (value.Type().Kind() == reflect.Struct ||
				value.Type().Kind() == reflect.Map ||
				value.Type().Kind() == reflect.Slice) {
				pType := reflect.New(value.Type())
				v2 := pType.Interface()
				err := json.Unmarshal([]byte(v.String()), &v2)
				if err == nil && v2 != nil {
					value.Set(reflect.ValueOf(v2).Elem())
				}
				// map, struct, slice to string
			} else if (field.Type.Kind() == reflect.Struct ||
				field.Type.Kind() == reflect.Map ||
				field.Type.Kind() == reflect.Slice) && value.Type().Kind() == reflect.String {
				ba, err := json.Marshal(v.Interface())
				if err == nil {
					val := string(ba)
					if strings.Contains(val, "{") {
						value.Set(reflect.ValueOf(string(ba)))
					} else {
						value.Set(reflect.ValueOf(""))
					}
				}
			} else if field.Type.Kind() != value.Type().Kind() { // 不同类型的字段过滤掉
				continue
			} else { // 简单数据类型的强制类型转换
				switch value.Kind() {
				case reflect.Int:
				case reflect.Int8:
				case reflect.Int16:
				case reflect.Int32:
				case reflect.Int64:
					value.SetInt(v.Int())
					break
				case reflect.Float32:
				case reflect.Float64:
					value.SetFloat(v.Float())
					break
				case reflect.Bool:
					value.SetBool(v.Bool())
					break
				}
			}

		}
	}

	return nil
}

func RandomNumber(bit int) int {
	minNum := intPow(10, bit-1)
	maxNum := intPow(10, bit) - 1

	rand.NewSource(time.Now().UnixNano())
	return rand.Intn(maxNum-minNum+1) + minNum
}

func intPow(x, y int) int {
	result := 1
	for i := 0; i < y; i++ {
		result *= x
	}
	return result
}
