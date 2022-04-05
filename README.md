iJSON
======

使用示例

```golang
  s := `{"object_key":{},"array_key":[],"字符串key":"字符串value","number_key_double":3.1415,` +
		`"number_key_int":9981,"bool_true_key":true,"bool_false_key":false,"null_key":null}`
  o := iJSON.ParseObject(iJSON.JSONString(s))
  log.Println(o.AsString())
  
  s = `[{},[],"字符串value",3.1415,9981,true,false,null]`
  a := iJSON.ParseArray(iJSON.JSONString(s))
  log.Println(a.AsString())
```
