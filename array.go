package iJSON

type JSONArray struct {
	array []*JSONValue
	count int
}

func NewJSONArray() *JSONArray {
	return &JSONArray{array: nil, count: 0}
}

func (a *JSONArray) Count() int {
	return a.count
}

func (a *JSONArray) getType() ValueType {
	return ValTypeArray
}

func (a *JSONArray) toBool() bool {
	return true
}

func (a *JSONArray) toDouble() float64 {
	return 0
}

func (a *JSONArray) toString(depth int, format bool) (s string) {
	if a.Count() < 1 {
		return `[]`
	}
	s += `[`
	depth++
	if format {
		s += "\n"
	}

	for idx, value := range a.array {
		if format {
			for i1 := 0; i1 < depth; i1++ {
				s += "\t"
			}
		}
		v := value.toString(depth, format)
		if value.ValueType == ValTypeString {
			s += `"` + v + `"`
		} else {
			s += v
		}

		if idx+1 < a.Count() {
			s += `,`
		}

		if format {
			s += "\n"
		}
	}

	if format {
		for i1 := 0; i1 < depth-1; i1++ {
			s += "\t"
		}
	}
	s += `]`
	return
}

func (a *JSONArray) AsString() string {
	return a.toString(0, false)
}

func (a *JSONArray) AsJSON() string {
	return a.toString(0, true)
}

func (a *JSONArray) Delete(index int) *JSONValue {
	if a.count < 1 || index >= a.count {
		return nil
	}
	v := a.array[index]
	a.array = append(a.array[:index], a.array[index+1:]...)
	a.count = len(a.array)
	return v
}

func (a *JSONArray) GetValue(index int) *JSONValue {
	if index >= a.Count() {
		return nil
	}
	return a.array[index]
}

func (a *JSONArray) AddValue(value *JSONValue) {
	if value == nil {
		value = newNullValue()
	}
	a.array = append(a.array, value)
	a.count = len(a.array)
}

func (a *JSONArray) GetObject(index int) *JSONObject {
	v := a.GetValue(index)
	return v.AsObject()
}

func (a *JSONArray) AddObject(value *JSONObject) {
	if value == nil {
		a.AddValue(nil)
	} else {
		a.AddValue(newObjectValue(value))
	}
}

func (a *JSONArray) GetArray(index int) *JSONArray {
	v := a.GetValue(index)
	return v.AsArray()
}

func (a *JSONArray) AddArray(value *JSONArray) {
	if value == nil {
		a.AddValue(nil)
	} else {
		a.AddValue(newArrayValue(value))
	}
}

func (a *JSONArray) GetString(index int) string {
	return a.GetValue(index).AsString()
}

func (a *JSONArray) AddString(value string) {
	a.AddValue(newStringValue(JSONString(value)))
}

func (a *JSONArray) GetDouble(index int) float64 {
	return a.GetValue(index).AsDouble()
}

func (a *JSONArray) AddDouble(value float64) {
	a.AddValue(newDoubleValue(value))
}

func (a *JSONArray) GetFloat(index int) float32 {
	return a.GetValue(index).AsFloat()
}

func (a *JSONArray) AddFloat(value float32) {
	a.AddValue(newFloatValue(value))
}

func (a *JSONArray) GetInt64(index int) int64 {
	return a.GetValue(index).AsInt64()
}

func (a *JSONArray) AddInt64(value int64) {
	a.AddValue(newIntValue(value))
}

func (a *JSONArray) GetInt(index int) int {
	return int(a.GetInt64(index))
}

func (a *JSONArray) AddInt(value int) {
	a.AddInt64(int64(value))
}

func (a *JSONArray) GetBool(index int) bool {
	return a.GetValue(index).AsBool()
}

func (a *JSONArray) AddBool(value bool) {
	if value {
		a.AddValue(newTrueValue())
	} else {
		a.AddValue(newFalseValue())
	}
}
