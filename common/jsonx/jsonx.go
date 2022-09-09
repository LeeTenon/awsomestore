// package jsonx 基于encoding/json实现，提供以下功能：
//	设置通过tag设置json unmarshal时的默认值
//	在marshal时初始化数组，使json中不会出现null
package jsonx

import (
	"encoding/json"
	"reflect"
	"strconv"
)

func Unmarshal(src []byte, dst interface{}) error {
	err := json.Unmarshal(src, dst)
	if err != nil {
		return err
	}

	typeof := reflect.TypeOf(dst)
	valueof := reflect.ValueOf(dst)
	for i := 0; i < typeof.Elem().NumField(); i++ {
		if valueof.Elem().Field(i).IsZero() {
			def := typeof.Elem().Field(i).Tag.Get("default")
			if def != "" {
				switch typeof.Elem().Field(i).Type.String() {
				case "int":
					result, _ := strconv.Atoi(def)
					valueof.Elem().Field(i).SetInt(int64(result))
				case "uint":
					result, _ := strconv.ParseUint(def, 10, 64)
					valueof.Elem().Field(i).SetUint(result)
				case "string":
					valueof.Elem().Field(i).SetString(def)
				}
			}
		}
	}

	return nil
}
