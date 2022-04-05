package iJSON

import (
	"testing"
)

func TestObjectParseSuccess(t *testing.T) {
	t.Log(`以下展示了JSON对象解析成功后的操作`)
	s := `{"杂key\"T\\T\/T\bT\fT\nT\rT\tT\u4e2d\u6587":null}`
	v := ParseValue(JSONString(s))
	if v == nil {
		t.Errorf("%s > JSON数据解析失败", s)
		return
	}
	o := v.AsObject()
	if o == nil {
		t.Errorf("%s > 非JSON对象数据", s)
	} else {
		t.Log(o.AsString())
	}

	s = `{"object_key":{},"array_key":[],"字符串key":"字符串value","number_key_double":3.1415,` +
		`"number_key_int":9981,"bool_true_key":true,"bool_false_key":false,"null_key":null}`
	if o = ParseObject(JSONString(s)); o == nil {
		t.Errorf("%s > JSON对象解析失败", s)
	} else {
		t.Log("\n", o.AsJSON())
		t.Logf("object_key=%s", o.GetObject("object_key").AsString())
		t.Logf("array_key=%s", o.GetArray("array_key").AsString())
		t.Logf("字符串key=%s", o.GetString("字符串key"))
		t.Logf("number_key_double=%.4f", o.GetDouble("number_key_double"))
		t.Logf("number_key_int=%d", o.GetInt("number_key_int"))
		t.Logf("bool_true_key=%t", o.GetBool("bool_true_key"))
		t.Logf("bool_false_key=%t", o.GetBool("bool_false_key"))
		t.Logf("null_key=%s", o.GetString("null_key"))
	}
}

func TestObjectParseFail(t *testing.T) {
	t.Log(`以下展示了对错误JSON对象数据的解析`)
	s := `{}{`
	o := ParseObject(JSONString(s))
	if o == nil {
		t.Errorf("%s > JSON对象解析失败", s)
	}

	s = `{"key":}`
	if o = ParseObject(JSONString(s)); o == nil {
		t.Errorf("%s > JSON对象解析失败", s)
	}

	s = `{"key":{}`
	if o = ParseObject(JSONString(s)); o == nil {
		t.Errorf("%s > JSON对象解析失败", s)
	}

	s = `{"key":}}`
	if o = ParseObject(JSONString(s)); o == nil {
		t.Errorf("%s > JSON对象解析失败", s)
	}

	s = `{"key":[}`
	if o = ParseObject(JSONString(s)); o == nil {
		t.Errorf("%s > JSON对象解析失败", s)
	}

	s = `{"key":]}`
	if o = ParseObject(JSONString(s)); o == nil {
		t.Errorf("%s > JSON对象解析失败", s)
	}
}

func TestObjectSetSuccess(t *testing.T) {
	t.Log(`以下展示了对JSON对象数据的生成`)
	o := NewJSONObject()
	//t.Log(o.AsString())

	o.SetObject(`对象`, NewJSONObject())
	//t.Log(o.AsString())

	o.SetArray(`数组`, NewJSONArray())
	//t.Log(o.AsString())

	o.SetString(`字符串`, `abc`)
	//t.Log(o.AsString())

	o.SetDouble(`double`, 3.1415926)
	//t.Log(o.AsString())

	o.SetFloat(`float`, 3.15)
	//t.Log(o.AsString())

	o.SetInt64(`int64`, 9876543210)
	//t.Log(o.AsString())

	o.SetInt(`int`, 65535)
	//t.Log(o.AsString())

	o.SetBool(`bool_true`, true)
	//t.Log(o.AsString())

	o.SetBool(`bool_false`, false)
	//t.Log(o.AsString())

	o.SetValue(`nil`, nil)
	//t.Log(o.AsString())

	t.Log("\n", o.AsJSON())
}
