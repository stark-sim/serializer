package serializer

import (
	"bytes"
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
)

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
		}
		for _, model := range models {
			var tempRes map[string]interface{}
			_json, _ := json.Marshal(model)
			decoder := json.NewDecoder(bytes.NewReader(_json))
			decoder.UseNumber()
			_ = decoder.Decode(&tempRes)
			for k, v := range tempRes {
				number, ok := v.(json.Number)
				// if the number is acceptable for frontend, make it back as a number, otherwise make it as a string
				if ok {
					intRes, err := number.Int64()
					if err != nil {
						if errors.Is(err, strconv.ErrSyntax) {
							floatRes, err := number.Float64()
							if err != nil {
								tempRes[k] = number.String()
							} else {
								if floatRes > float64(MinSafeNumber) && floatRes < float64(MaxSafeNumber) {
									tempRes[k] = floatRes
								} else {
									tempRes[k] = number.String()
								}
							}
						} else {
							tempRes[k] = number.String()
						}
					} else {
						if intRes > MinSafeNumber && intRes < MaxSafeNumber {
							tempRes[k] = intRes
						} else {
							tempRes[k] = number.String()
						}
					}
				}
			}
			resList = append(resList, tempRes)
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
