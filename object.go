package iJSON

type JSONObject struct {
	tree *AvlTree
}

func NewJSONObject() *JSONObject {
	return &JSONObject{tree: NewAvlTree()}
}

func (o *JSONObject) getType() ValueType {
	return ValTypeObject
}

func (o *JSONObject) toBool() bool {
	return true
}

func (o *JSONObject) toDouble() float64 {
	return 0
}

func (o *JSONObject) toString(depth int, format bool) (s string) {
	if o.tree.Count() < 1 {
		return `{}`
	}
	s += `{`
	depth++
	if format {
		s += "\n"
	}
	keys := o.tree.GetKeys()
	for idx, key := range keys {
		value := o.GetValue(key)
		if format {
			for i1 := 0; i1 < depth; i1++ {
				s += "\t"
			}
		}
		s += `"` + key + `"`
		s += `:`
		if format {
			s += ` `
		}
		v := value.toString(depth, format)
		if value.ValueType == ValTypeString {
			s += `"` + v + `"`
		} else {
			s += v
		}

		if idx+1 < o.tree.Count() {
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
	s += `}`
	return
}

func (o *JSONObject) AsString() string {
	return o.toString(0, false)
}

func (o *JSONObject) AsJSON() string {
	return o.toString(0, true)
}

func (o *JSONObject) Delete(key string) *JSONValue {
	an := o.tree.Remove(key)
	if an == nil {
		return nil
	}
	return an.value.(*JSONValue)
}

func (o *JSONObject) GetValue(key string) *JSONValue {
	an := o.tree.Get(key)
	if an == nil {
		return nil
	}
	return an.value.(*JSONValue)
}

func (o *JSONObject) SetValue(key string, value *JSONValue) {
	if value == nil {
		value = newNullValue()
	}
	o.tree.Set(key, value)
}

func (o *JSONObject) GetObject(key string) *JSONObject {
	return o.GetValue(key).AsObject()
}

func (o *JSONObject) SetObject(key string, value *JSONObject) {
	if value == nil {
		o.SetValue(key, nil)
	} else {
		o.SetValue(key, newObjectValue(value))
	}
}

func (o *JSONObject) GetArray(key string) *JSONArray {
	return o.GetValue(key).AsArray()
}

func (o *JSONObject) SetArray(key string, value *JSONArray) {
	if value == nil {
		o.SetValue(key, nil)
	} else {
		o.SetValue(key, newArrayValue(value))
	}
}

func (o *JSONObject) GetString(key string) string {
	return o.GetValue(key).AsString()
}

func (o *JSONObject) SetString(key, value string) {
	o.SetValue(key, newStringValue(JSONString(value)))
}

func (o *JSONObject) GetDouble(key string) float64 {
	return o.GetValue(key).AsDouble()
}

func (o *JSONObject) SetDouble(key string, value float64) {
	o.SetValue(key, newDoubleValue(value))
}

func (o *JSONObject) GetFloat(key string) float32 {
	return o.GetValue(key).AsFloat()
}

func (o *JSONObject) SetFloat(key string, value float32) {
	o.SetValue(key, newFloatValue(value))
}

func (o *JSONObject) GetInt64(key string) int64 {
	return o.GetValue(key).AsInt64()
}

func (o *JSONObject) SetInt64(key string, value int64) {
	o.SetValue(key, newIntValue(value))
}

func (o *JSONObject) GetInt(key string) int {
	return int(o.GetInt64(key))
}

func (o *JSONObject) SetInt(key string, value int) {
	o.SetInt64(key, int64(value))
}

func (o *JSONObject) GetBool(key string) bool {
	return o.GetValue(key).AsBool()
}

func (o *JSONObject) SetBool(key string, value bool) {
	if value {
		o.SetValue(key, newTrueValue())
	} else {
		o.SetValue(key, newFalseValue())
	}
}
