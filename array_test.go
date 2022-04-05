package iJSON

import "testing"

func TestArrayParseSuccess(t *testing.T) {
	t.Log(`以下展示了JSON数组解析成功后的操作`)
	s := `["\"T\\T\/T\bT\fT\nT\rT\tT\u4e2d\u6587"]`
	v := ParseValue(JSONString(s))
	if v == nil {
		t.Errorf("%s > JSON数据解析失败", s)
		return
	}
	a := v.AsArray()
	if a == nil {
		t.Errorf("%s > 非JSON数组数据", s)
	} else {
		t.Log(a.AsString())
	}

	s = `[{},[],"字符串value",3.1415,9981,true,false,null]`
	if a = ParseArray(JSONString(s)); a == nil {
		t.Errorf("%s > JSON数组解析失败", s)
	} else {
		t.Log("\n", a.AsJSON())
		t.Logf("序列 0=%s", a.GetObject(0).AsString())
		t.Logf("序列 1=%s", a.GetArray(1).AsString())
		t.Logf("序列 2=%s", a.GetString(2))
		t.Logf("序列 3=%.4f", a.GetDouble(3))
		t.Logf("序列 4=%d", a.GetInt(4))
		t.Logf("序列 5=%t", a.GetBool(5))
		t.Logf("序列 6=%t", a.GetBool(6))
		t.Logf("序列 7=%s", a.GetString(7))
	}
}

func TestArrayParseFail(t *testing.T) {
	t.Log(`以下展示了对错误JSON数组数据的解析`)
	s := `[][`
	o := ParseObject(JSONString(s))
	if o == nil {
		t.Errorf("%s > JSON数组解析失败", s)
	}

	s = `[]]`
	if o = ParseObject(JSONString(s)); o == nil {
		t.Errorf("%s > JSON数组解析失败", s)
	}

	s = `[{]`
	if o = ParseObject(JSONString(s)); o == nil {
		t.Errorf("%s > JSON数组解析失败", s)
	}

	s = `[}]`
	if o = ParseObject(JSONString(s)); o == nil {
		t.Errorf("%s > JSON数组解析失败", s)
	}

	s = `["a]`
	if o = ParseObject(JSONString(s)); o == nil {
		t.Errorf("%s > JSON数组解析失败", s)
	}

	s = `[b"]`
	if o = ParseObject(JSONString(s)); o == nil {
		t.Errorf("%s > JSON数组解析失败", s)
	}
}

func TestArrayAddSuccess(t *testing.T) {
	t.Log(`以下展示了对JSON数组数据的生成`)
	a := NewJSONArray()
	a.AddObject(NewJSONObject())
	a.AddArray(NewJSONArray())
	a.AddString("字符串")
	a.AddDouble(3.1415926)
	a.AddFloat(3.15)
	a.AddInt64(9876543210)
	a.AddInt(65535)
	a.AddBool(true)
	a.AddBool(false)
	a.AddValue(nil)

	t.Log("\n", a.AsJSON())
}
