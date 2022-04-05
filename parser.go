package iJSON

import "strconv"

type parser struct {
	tok   *tokenizer //分词器
	count int        //词列表长度
	index int        //读取标记所在位置
}

func newParser() *parser {
	return &parser{}
}

func (o *parser) init(s JSONString) {
	o.tok = newTokenizer(s)
	o.tok.DoExec()
	o.count = len(o.tok.termArray)
	o.index = -1
}

func (o *parser) next() bool {
	o.index++
	return o.index < o.count
}

func (o *parser) read() *term {
	return o.tok.termArray[o.index]
}

func (o *parser) parseKey() JSONString {
	if !o.next() {
		return nil
	}
	//期望string
	if tm := o.read(); tm.tp != TmString {
		return nil
	} else {
		return tm.value
	}
}

func (o *parser) parseValue() (rsVal *JSONValue) {
	if !o.next() {
		return nil
	}

	rsVal = new(JSONValue)
	tm := o.read()
	switch tm.tp {
	case TmObjectBegin:
		rsVal = newObjectValue(o.parseObject())
	case TmArrayBegin:
		rsVal = newArrayValue(o.parseArray())
	case TmString:
		rsVal = newStringValue(tm.value)
	case TmNumber:
		rsVal = o.parseNumber(tm.value)
	case TmTrue:
		rsVal = newTrueValue()
	case TmFalse:
		rsVal = newFalseValue()
	case TmNull:
		rsVal = newNullValue()
	default:
		rsVal = nil
	}
	return
}

func (o *parser) parseObject() (rsObj *JSONObject) {
	var tm *term
	if !o.next() {
		return nil
	}
	rsObj = NewJSONObject()
	if tm = o.read(); tm.tp == TmObjectEnd {
		return
	}
	o.index--
	getOne := func() bool {
		k := o.parseKey()
		if k == nil || !o.next() {
			return false
		}
		//期望冒号
		if tm = o.read(); tm.tp != TmSepColon {
			return false
		}
		v := o.parseValue()
		if v == nil {
			return false
		}
		rsObj.SetValue(k.AsString(), v)
		return true
	}

	if !getOne() {
		return nil
	}

	for o.next() {
		tm = o.read()
		if tm.tp != TmSepComma {
			break
		}

		if !getOne() {
			return nil
		}
	}

	if tm.tp != TmObjectEnd {
		return nil
	}
	return rsObj
}

func (o *parser) parseArray() (rsArr *JSONArray) {
	var tm *term
	if !o.next() {
		return nil
	}
	rsArr = NewJSONArray()
	if tm = o.read(); tm.tp == TmArrayEnd {
		return
	}
	o.index--
	v := o.parseValue()
	if v == nil {
		return nil
	}
	rsArr.AddValue(v)

	for o.next() {
		tm = o.read()
		if tm.tp != TmSepComma {
			break
		}

		if v = o.parseValue(); v == nil {
			return nil
		}
		rsArr.AddValue(v)
	}

	if tm.tp != TmArrayEnd {
		return nil
	}
	return rsArr
}

func (o *parser) parseNumber(value JSONString) *JSONValue {
	for _, c := range value {
		switch c {
		case '.', 'e', 'E':
			d, _ := strconv.ParseFloat(string(value), 64)
			return newDoubleValue(d)
		}
	}
	i, _ := strconv.ParseInt(string(value), 10, 64)
	return newIntValue(i)
}

func ParseValue(s JSONString) (rsVal *JSONValue) {
	o := newParser()
	if o.init(s); o.tok.err != nil {
		return nil
	}
	rsVal = o.parseValue()
	if o.next() { //存在后续词，JSON数据可能有误，如{}{或{}[]
		return nil
	}
	return rsVal
}

func ParseObject(s JSONString) (rsObj *JSONObject) {
	v := ParseValue(s)
	if v == nil {
		return nil
	}
	rsObj, _ = v.val.(*JSONObject)
	return rsObj
}

func ParseArray(s JSONString) (rsArr *JSONArray) {
	v := ParseValue(s)
	if v == nil {
		return nil
	}
	rsArr, _ = v.val.(*JSONArray)
	return rsArr
}
