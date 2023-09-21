package serializer

import (
	"bytes"
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
	"strings"
)

const (
	MinSafeNumber int64 = -1 << 53
	MaxSafeNumber int64 = 1<<53 - 1
)

func FrontendSerialize(model interface{}) interface{} {
	jsonBytes, _ := json.Marshal(model)
	// json decode once
	decoder := json.NewDecoder(bytes.NewReader(jsonBytes))
	decoder.UseNumber()
	var res interface{}
	_ = decoder.Decode(&res)
	// data adjustment O(n)
	return adjustSerializedObj(res)
}

// 调整序列化后的数据，应在数据已经 json 化后进行调用，即只包含 []map[string]interface{}, map[string]interface{}, json.number 和 基础类型
func adjustSerializedObj(value interface{}) interface{} {
	rv := reflect.ValueOf(value)
	// 如果是对象或是列表，则有可能存在复杂数据，需要再次递归
	if rv.Kind() == reflect.Slice {
		// https://stackoverflow.com/questions/52479739/how-to-convert-interface-to-interfaces-slice
		// interface{} to []interface{}, this solution is bravo!
		var values []interface{}
		for i := 0; i < rv.Len(); i++ {
			values = append(values, adjustSerializedObj(rv.Index(i).Interface()))
		}
		return values
	} else if obj, ok := value.(map[string]interface{}); ok {
		for k, v := range obj {
			// 如果 key 是 id 结尾的，那直接转 string
			if strings.HasSuffix(k, "id") {
				strID, ok := v.(json.Number)
				if ok {
					obj[k] = strID.String()
				} else {
					obj[k] = adjustSerializedObj(v)
				}
			} else {
				obj[k] = adjustSerializedObj(v)
			}
		}
		return obj
		// 	如果是数字，检查是否可以转换回数字类型，而非全定义为字符串
	} else if number, ok := value.(json.Number); ok {
		intRes, err := number.Int64()
		if err != nil {
			if errors.Is(err, strconv.ErrSyntax) {
				floatRes, err := number.Float64()
				if err != nil {
					return number.String()
				} else {
					return floatRes
				}
			} else {
				return number.String()
			}
		} else {
			if intRes > MinSafeNumber && intRes < MaxSafeNumber {
				return intRes
			} else {
				return number.String()
			}
		}
	} else {
		return value
	}
}
