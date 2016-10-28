package jvalue

import (
	"encoding/json"
	"errors"
)

// JValue provides script language style json access
type JValue struct {
	body interface{}
}

var emptyJValue = JValue{body: nil}

func (jValue *JValue) ToString() (*string, error) {
	value, flag := jValue.body.(string)
	if flag == false {
		return nil, errors.New("cant parse")
	}
	return &value, nil
}
func (jValue *JValue) ToInt() (*int, error) {
	value, flag := jValue.body.(float64)
	if flag == false {
		return nil, errors.New("cant parse")
	}
	number := int(value)
	return &number, nil
}
func (jValue *JValue) ToBool() (*bool, error) {
	value, flag := jValue.body.(bool)
	if flag == false {
		return nil, errors.New("cant parse")
	}
	return &value, nil
}

// DecodeJSONString return JValue.
func DecodeJSONString(jsonStr string) (JValue, error) {
	var data interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return emptyJValue, err
	}
	return JValue{body: data}, nil
}

func (jValue *JValue) IsArray() bool {
	if jValue == nil || jValue.body == nil {
		return false
	}
	_, isArr := jValue.body.([]interface{})
	return isArr
}
func (jValue *JValue) HasIndex(index int) bool {
	if !jValue.IsArray() {
		return false
	}
	jArr := jValue.body.([]interface{})
	return len(jArr) > index
}
func (jValue *JValue) Index(index int) *JValue {
	if !jValue.IsArray() || !jValue.HasIndex(index) {
		return &JValue{body: nil}
	}
	jArr := jValue.body.([]interface{})
	return &JValue{body: jArr[index]}
}

func (jValue *JValue) IsMap() bool {
	_, isMap := jValue.body.(map[string]interface{})
	return isMap
}

func (jValue *JValue) HasKey(key string) bool {
	if !jValue.IsMap() {
		return false
	}
	jMap := jValue.body.(map[string]interface{})
	_, flag := jMap[key]
	return flag
}

func (jValue *JValue) Key(key string) *JValue {
	if !jValue.IsMap() || !jValue.HasKey(key) {
		return &JValue{body: nil}
	}
	jMap := jValue.body.(map[string]interface{})
	return &JValue{body: jMap[key]}
}
