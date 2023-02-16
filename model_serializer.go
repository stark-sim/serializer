package serializer

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// ModelSerializer 序列化器，需要手动指定是否为 slice，但优点是性能更好一点
type ModelSerializer struct {
	Model    interface{}
	IsPlural bool
}

const (
	MinSafeNumber int64 = -1 << 53
	MaxSafeNumber int64 = 1<<53 - 1
)

func (s *ModelSerializer) Serialize() interface{} {
	if s.IsPlural {
		var resList []map[string]interface{}
		var models []interface{}
		// https://stackoverflow.com/questions/52479739/how-to-convert-interface-to-interfaces-slice
		// interface{} to []interface{}, this solution is bravo!
		rv := reflect.ValueOf(s.Model)
		if rv.Kind() == reflect.Slice {
			for i := 0; i < rv.Len(); i++ {
				models = append(models, rv.Index(i).Interface())
			}
		} else {
			return nil
		}
		_json, _ := json.Marshal(models)
		decoder := json.NewDecoder(bytes.NewReader(_json))
		decoder.UseNumber()
		_ = decoder.Decode(&resList)
		for _, item := range models {
			rv := reflect.ValueOf(item)
			fmt.Println(rv.Kind())
		}
		return resList
	} else {
		var res map[string]interface{}
		_json, _ := json.Marshal(s.Model)
		decoder := json.NewDecoder(bytes.NewReader(_json))
		decoder.UseNumber()
		_ = decoder.Decode(&res)
		for k, v := range res {
			number, ok := v.(json.Number)
			// if the number is acceptable for frontend, make it back as a number, otherwise make it as a string
			if ok {
				intRes, err := number.Int64()
				if err != nil {
					if errors.Is(err, strconv.ErrSyntax) {
						floatRes, err := number.Float64()
						if err != nil {
							res[k] = number.String()
						} else {
							if floatRes > float64(MinSafeNumber) && floatRes < float64(MaxSafeNumber) {
								res[k] = floatRes
							} else {
								res[k] = number.String()
							}
						}
					} else {
						res[k] = number.String()
					}
				} else {
					if intRes > MinSafeNumber && intRes < MaxSafeNumber {
						res[k] = intRes
					} else {
						res[k] = number.String()
					}
				}
			}
		}
		return res
	}
}

func adjustSerializedObj(value map[string]interface{}) map[string]interface{} {
	for k, v := range value {
		number, ok := v.(json.Number)
		if ok {
			intRes, err := number.Int64()
			if err != nil {
				if errors.Is(err, strconv.ErrSyntax) {
					floatRes, err := number.Float64()
					if err != nil {
						value[k] = number.String()
					} else {
						value[k] = floatRes
					}
				} else {
					value[k] = number.String()
				}
			} else {
				if intRes > MinSafeNumber && intRes < MaxSafeNumber {
					value[k] = intRes
				} else {
					value[k] = number.String()
				}
			}
		} else {
			obj, ok := v.(map[string]interface{})
			if ok {
				value[k] = adjustSerializedObj(obj)
			} else {
				list, ok := v.([]interface{})
				if ok {
					var tempRes []map[string]interface{}
					for _, item := range list {
						// TODO deal with [snowflakeID_0, snowflakeID_1]
						tempRes = append(tempRes, adjustSerializedObj(item.(map[string]interface{})))
					}
					value[k] = tempRes
				}
			}
		}
	}
	return value
}

func Serialize(model interface{}) interface{} {
	rv := reflect.ValueOf(model)
	// https://stackoverflow.com/questions/52479739/how-to-convert-interface-to-interfaces-slice
	// interface{} to []interface{}, this solution is bravo!
	if rv.Kind() == reflect.Slice {
		var models []interface{}
		for i := 0; i < rv.Len(); i++ {
			// TODO 试着直接序列化
			models = append(models, Serialize(rv.Index(i).Interface()))
		}
		return models
	} else {
		var res map[string]interface{}
		// model to json, number all become string
		_json, _ := json.Marshal(model)
		decoder := json.NewDecoder(bytes.NewReader(_json))
		decoder.UseNumber()
		_ = decoder.Decode(&res)
		// check maybe some string can turn back to number
		res = adjustSerializedObj(res)
		return res
	}
}
