package iJSON

import "strconv"

type ValueType int

const (
	ValTypeNULL   ValueType = 100 //NULL
	ValTypeObject ValueType = 101 //对象型
	ValTypeArray  ValueType = 102 //数组型
	ValTypeString ValueType = 103 //字符型
	ValTypeTrue   ValueType = 104 //True
	ValTypeFalse  ValueType = 105 //False
	ValTypeDouble ValueType = 106 //双精度
	ValTypeFloat  ValueType = 107 //单精度
	ValTypeInt    ValueType = 108 //整型
)

// JSONValue //////////////////////////////////////////////////////////////////////////////////////////////////
type JSONValue struct {
	ValueType
	val interface{}
}

func newObjectValue(value *JSONObject) *JSONValue {
	if value == nil {
		return nil
	}

	return &JSONValue{
		ValueType: ValTypeObject,
		val:       value,
	}
}

func newArrayValue(value *JSONArray) *JSONValue {
	if value == nil {
		return nil
	}

	return &JSONValue{
		ValueType: ValTypeArray,
		val:       value,
	}
}

func newStringValue(value JSONString) *JSONValue {
	return &JSONValue{
		ValueType: ValTypeString,
		val:       value,
	}
}

func newDoubleValue(value float64) *JSONValue {
	return &JSONValue{
		ValueType: ValTypeDouble,
		val:       value,
	}
}

func newFloatValue(value float32) *JSONValue {
	return &JSONValue{
		ValueType: ValTypeFloat,
		val:       value,
	}
}

func newIntValue(value int64) *JSONValue {
	return &JSONValue{
		ValueType: ValTypeInt,
		val:       value,
	}
}

func newTrueValue() *JSONValue {
	return &JSONValue{
		ValueType: ValTypeTrue,
		val:       true,
	}
}

func newFalseValue() *JSONValue {
	return &JSONValue{
		ValueType: ValTypeFalse,
		val:       false,
	}
}

func newNullValue() *JSONValue {
	return &JSONValue{
		ValueType: ValTypeNULL,
		val:       nil,
	}
}

func (v *JSONValue) AsDouble() float64 {
	if v == nil {
		return 0
	}

	switch v.ValueType {
	case ValTypeDouble:
		return v.val.(float64)
	case ValTypeFloat:
		return float64(v.val.(float32))
	case ValTypeInt:
		return float64(v.val.(int64))
	default:
		return 0
	}
}

func (v *JSONValue) AsFloat() float32 {
	return float32(v.AsDouble())
}

func (v *JSONValue) AsInt64() int64 {
	if v == nil {
		return 0
	}

	switch v.ValueType {
	case ValTypeDouble:
		return int64(v.val.(float64))
	case ValTypeFloat:
		return int64(v.val.(float32))
	case ValTypeInt:
		return v.val.(int64)
	default:
		return 0
	}
}

func (v *JSONValue) AsInt() int {
	return int(v.AsInt64())
}

func (v *JSONValue) AsBool() bool {
	if v == nil {
		return false
	}

	switch v.ValueType {
	case ValTypeDouble:
		return v.val.(float64) != 0
	case ValTypeFloat:
		return v.val.(float32) != 0
	case ValTypeInt:
		return v.val.(int64) != 0
	case ValTypeString:
		return len(v.val.(JSONString)) > 0
	case ValTypeObject:
		return v.val.(*JSONObject) != nil
	case ValTypeArray:
		return v.val.(*JSONArray) != nil
	case ValTypeNULL, ValTypeFalse:
		return false
	case ValTypeTrue:
		return true
	default:
		return false
	}
}

func (v *JSONValue) toString(depth int, format bool) string {
	if v == nil {
		return ``
	}

	switch v.ValueType {
	case ValTypeObject:
		return v.val.(*JSONObject).toString(depth, format)
	case ValTypeArray:
		return v.val.(*JSONArray).toString(depth, format)
	case ValTypeString:
		return v.val.(JSONString).AsString()
	case ValTypeDouble:
		return FormatDouble(v.val.(float64))
	case ValTypeFloat:
		return FormatFloat(v.val.(float32))
	case ValTypeInt:
		return strconv.FormatInt(v.val.(int64), 10)
	case ValTypeTrue:
		return `true`
	case ValTypeFalse:
		return `false`
	case ValTypeNULL:
		return `null`
	default:
		return ``
	}
}

func (v *JSONValue) AsJSON() string {
	return v.toString(0, true)
}

func (v *JSONValue) AsString() string {
	return v.toString(0, false)
}

func (v *JSONValue) AsObject() *JSONObject {
	if v == nil {
		return nil
	}

	switch v.ValueType {
	case ValTypeObject:
		return v.val.(*JSONObject)
	default:
		return nil
	}
}

func (v *JSONValue) AsArray() *JSONArray {
	if v == nil {
		return nil
	}

	switch v.ValueType {
	case ValTypeArray:
		return v.val.(*JSONArray)
	default:
		return nil
	}
}

// JSONString //////////////////////////////////////////////////////////////////////////////////////////////////
type JSONString []rune

//func (s JSONString) toString() string {
//	return s.AsString()
//}

func (s JSONString) getType() ValueType {
	return ValTypeString
}

func (s JSONString) AsString() string {
	var sr, s1 string
	for _, c := range s {
		switch c {
		case '"':
			s1 = `\"`
		case '\\':
			s1 = `\\`
		case '/':
			s1 = `\/`
		case '\b':
			s1 = `\b`
		case '\f':
			s1 = `\f`
		case '\n':
			s1 = `\n`
		case '\r':
			s1 = `\r`
		case '\t':
			s1 = `\t`
		default:
			s1 = string(c)
		}
		sr += s1
	}
	return sr
}
