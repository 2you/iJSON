package iJSON

import (
	"testing"
	"time"
)

func TestObjectParse1(t *testing.T) {
	s := `{"\"\\\/\b\f\n\r\t\u4e2d\u6587":null,"arr_1":[8192,32],"arr_0":[],"obj_1":{"float":3.15},"bool_true":  true,` +
		`"圆周率":3.1415926,"obj_0":{},"bool_false":false,"当前时间":"2022-04-02 00:26:06","unix_time":1648830366977590300}`
	v := ParseValue(JSONString(s))
	if v == nil {
		return
	}
	t.Log("\n", v.AsJSON())
	o := v.AsObject()
	t.Log(o.GetString(`当前时间`))
	t.Log(o.GetInt64(`unix_time`))
	t.Log(o.GetDouble(`圆周率`))
	t.Log(o.GetBool(`bool_true`))
	t.Log("\n", o.AsJSON())

	s = `{"object_key":{},"array_key":[],"字符串key":"字符串value","number_key_double":3.1415,"number_key_int":9981,"bool_true_key":true,"bool_false_key":false,"null_key":null}`
	if o = ParseObject(JSONString(s)); o == nil {
		return
	}
	t.Log("\n", o.AsJSON())
	t.Log("object_key", o.GetObject("object_key").AsString())
	t.Log("array_key", o.GetArray("array_key").AsString())
	t.Log("字符串key", o.GetString("字符串key"))
	t.Log("number_key_double", o.GetDouble("number_key_double"))
	t.Log("number_key_int", o.GetInt("number_key_int"))
	t.Log("bool_true_key", o.GetBool("bool_true_key"))
	t.Log("bool_false_key", o.GetBool("bool_false_key"))
	t.Log("null_key", o.GetString("null_key"))
}

func TestObjectParse2(t *testing.T) {
	s := `{"\u0075\u006e\u0069\u0063\u006f\u0064\u0065\u4e2d\u6587\u4e92\u8f6c":true}`
	o := ParseObject(JSONString(s))
	t.Log(o.AsString())

	s = `{"k0":{"k1":{"k2":}}}`
	if o = ParseObject(JSONString(s)); o == nil {
		t.Errorf("%s >解析出错", s)
	} else {
		t.Log(o.AsString())
	}

	v := ParseValue(JSONString(s))
	t.Log(v.AsString())
}

func TestObjectSet(t *testing.T) {
	o := NewJSONObject()
	o.SetString(`当前时间`, time.Now().Format(`2006-01-02 15:04:05`))
	o.SetInt64(`unix_time`, time.Now().UnixNano())
	o.SetDouble(`圆周率`, 3.1415926)
	o.SetBool(`bool_true`, true)
	o.SetBool(`bool_false`, false)
	o.SetObject(`nil`, nil)
	o0 := NewJSONObject()
	o.SetObject(`obj_0`, o0)
	o1 := NewJSONObject()
	o.SetObject(`obj_1`, o1)
	o1.SetFloat(`float`, 3.15)
	a0 := NewJSONArray()
	o.SetArray(`arr_0`, a0)
	a1 := NewJSONArray()
	o.SetArray(`arr_1`, a1)
	a1.AddInt64(8192)
	a1.AddInt(32)
	t.Log(o1.AsString())
	t.Log(a1.AsString())
	t.Log(o.AsString())
	t.Log("\n", o.AsJSON())
}

func TestArrayParse(t *testing.T) {
	s := `[{},[],"字符串",3.1415,9981,true,false,null]`
	v := ParseValue(JSONString(s))
	t.Log("\n", v.AsJSON())

	a := ParseArray(JSONString(s))
	t.Log("\n", a.AsJSON())
}

func TestArraySet(t *testing.T) {
	a := NewJSONArray()
	a.AddString(`"\/` + "\n\r\t\b\f")
	t.Log("\n", a.AsJSON())
}
