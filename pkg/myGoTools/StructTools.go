package myGoTools

import (
	"fmt"
	"reflect"
	"strings"
)

// 获取struct的字段名称，返回字段切片
func GetFieldName(s interface{}) ([]string, error) {
	t := reflect.TypeOf(s)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("check type error not Struct")
	}
	fieldNum := t.NumField()
	result := make([]string, 0 , fieldNum)
	for i := 0; i < fieldNum; i++ {
		result = append(result, t.Field(i).Name)
	}
	return result, nil
}

// 获取结构体中的Tag的值，没有tag则返回字段值
func GetTagName(s interface{}) ([]string, error) {
	t := reflect.TypeOf(s)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("check type error not Struct")
	}
	fieldNum := t.NumField()
	result := make([]string,0,fieldNum)
	for i := 0; i < fieldNum; i++ {
		tagName := t.Field(i).Name
		tags := strings.Split(string(t.Field(i).Tag), "\"")
		if len(tags) > 1 {
			tagName = tags[1]
		}
		result = append(result,tagName)
	}
	return result, nil
}