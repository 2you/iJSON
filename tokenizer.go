package iJSON

import (
	"errors"
	"fmt"
	"strconv"
)

type TermType int

//词类型
const (
	TmObjectBegin TermType = 1  // {
	TmObjectEnd   TermType = 2  // }
	TmArrayBegin  TermType = 3  // [
	TmArrayEnd    TermType = 4  // ]
	TmString      TermType = 5  // "
	TmNumber      TermType = 6  // 数值
	TmTrue        TermType = 7  // true（不区分大小写）
	TmFalse       TermType = 8  // false（不区分大小写）
	TmNull        TermType = 9  // null（不区分大小写）
	TmSepColon    TermType = 10 // 冒号
	TmSepComma    TermType = 11 // 逗号
)

type term struct {
	tp    TermType   //词类型
	value JSONString //词值
}

func newTermA(tokenType TermType, value JSONString) *term {
	return &term{
		tp:    tokenType,
		value: value,
	}
}

func newTermB(tokenType TermType) *term {
	switch tokenType {
	case TmObjectBegin:
		return newTermA(TmObjectBegin, JSONString(`{`))
	case TmObjectEnd:
		return newTermA(TmObjectEnd, JSONString(`}`))
	case TmArrayBegin:
		return newTermA(TmArrayBegin, JSONString(`[`))
	case TmArrayEnd:
		return newTermA(TmArrayEnd, JSONString(`]`))
	case TmSepColon:
		return newTermA(TmSepColon, JSONString(`:`))
	case TmSepComma:
		return newTermA(TmSepComma, JSONString(`,`))
	default:
		return nil
	}
}

type tokenizer struct {
	data      JSONString //JSON数据
	size      int        //数据长度
	offset    int        //读取标记所在位置
	err       error      //执行过程中产生的错误
	termArray []*term    //获取到的词列表
}

func newTokenizer(data JSONString) (tok *tokenizer) {
	tok = new(tokenizer)
	tok.data = data
	tok.size = len(data)
	tok.offset = 0
	tok.err = nil
	return tok
}

func (tok *tokenizer) DoExec() {
	var c rune
	var s JSONString
	var tm *term
	for tok.offset < tok.size {
		if c, tok.err = tok.readChar(); tok.err != nil {
			return
		}
		switch c {
		case ' ':
			continue
		case '{':
			tm = newTermB(TmObjectBegin)
		case '}':
			tm = newTermB(TmObjectEnd)
		case '[':
			tm = newTermB(TmArrayBegin)
		case ']':
			tm = newTermB(TmArrayEnd)
		case ':':
			tm = newTermB(TmSepColon)
		case ',':
			tm = newTermB(TmSepComma)
		case '"':
			if s, tok.err = tok.readString(); tok.err == nil {
				tm = newTermA(TmString, s)
			} else {
				return
			}
		case 'n', 'N':
			if s, tok.err = tok.readNull(); tok.err == nil {
				s[0] = c
				tm = newTermA(TmNull, s)
			} else {
				return
			}
		case 't', 'T':
			if s, tok.err = tok.readTrue(); tok.err == nil {
				s[0] = c
				tm = newTermA(TmTrue, s)
			} else {
				return
			}
		case 'f', 'F':
			if s, tok.err = tok.readFalse(); tok.err == nil {
				s[0] = c
				tm = newTermA(TmFalse, s)
			} else {
				return
			}
		default: //数值型
			if s, tok.err = tok.readNumber(); tok.err == nil {
				s[0] = c
				if _, tok.err = strconv.ParseFloat(s.AsString(), 64); tok.err != nil {
					return
				}
				tm = newTermA(TmNumber, s)
			} else {
				return
			}
		}
		//log.Println(tm, tm.tp, tm.value.AsString())
		//log.Println(tm.value.AsString())
		tok.addTerm(tm)
	}
}

func (tok *tokenizer) addTerm(token *term) {
	tok.termArray = append(tok.termArray, token)
}

func (tok *tokenizer) readNull() (s JSONString, e error) {
	s = make(JSONString, 4)
	isErr := false
	var c rune
	for i1 := 1; i1 < 4; i1++ {
		if c, e = tok.readChar(); e != nil {
			return nil, e
		}

		if i1 == 1 && c != 'u' && c != 'U' {
			isErr = true
		} else if i1 == 2 && c != 'l' && c != 'L' {
			isErr = true
		} else if i1 == 3 && c != 'l' && c != 'L' {
			isErr = true
		}

		if isErr {
			return nil, fmt.Errorf("char index %d is not null type", tok.offset)
		}
		s[i1] = c
	}
	return s, e
}

func (tok *tokenizer) readTrue() (s JSONString, e error) {
	s = make(JSONString, 4)
	isErr := false
	var c rune
	for i1 := 1; i1 < 4; i1++ {
		if c, e = tok.readChar(); e != nil {
			return nil, e
		}

		if i1 == 1 && c != 'r' && c != 'R' {
			isErr = true
		} else if i1 == 2 && c != 'u' && c != 'U' {
			isErr = true
		} else if i1 == 3 && c != 'e' && c != 'E' {
			isErr = true
		}

		if isErr {
			return nil, fmt.Errorf("char index %d is not boolean true type", tok.offset)
		}
		s[i1] = c
	}
	return s, e
}

func (tok *tokenizer) readFalse() (s JSONString, e error) {
	s = make(JSONString, 5)
	isErr := false
	var c rune
	for i1 := 1; i1 < 5; i1++ {
		if c, e = tok.readChar(); e != nil {
			return nil, e
		}

		if i1 == 1 && c != 'a' && c != 'A' {
			isErr = true
		} else if i1 == 2 && c != 'l' && c != 'L' {
			isErr = true
		} else if i1 == 3 && c != 's' && c != 'S' {
			isErr = true
		} else if i1 == 4 && c != 'e' && c != 'E' {
			isErr = true
		}

		if isErr {
			return nil, fmt.Errorf("char index %d is not boolean false type", tok.offset)
		}
		s[i1] = c
	}
	return s, e
}

func (tok *tokenizer) readNumber() (s JSONString, e error) {
	var c rune
	s = make(JSONString, 1)
	for {
		if c, e = tok.readChar(); e != nil {
			return nil, e
		}

		if !((c >= '0' && c <= '9') || c == '.' || c == '+' ||
			c == '-' || c == 'e' || c == 'E') {
			tok.offset--
			break
		}
		s = append(s, c)
	}
	return s, nil
}

func (tok *tokenizer) readString() (s JSONString, e error) {
	var c rune
	var u int64
	rn := make(JSONString, 4)
	for {
		if c, e = tok.readChar(); e != nil {
			return nil, e
		}

		if c == '"' {
			break
		}

		if c == '\\' {
			if c, e = tok.readChar(); e != nil {
				return nil, e
			}

			switch c {
			case '"', '\\', '/':

			case 'b':
				c = '\b'
			case 'f':
				c = '\f'
			case 'n':
				c = '\n'
			case 'r':
				c = '\r'
			case 't':
				c = '\t'
			case 'u':
				for i2 := 0; i2 < 4; i2++ {
					if c, e = tok.readChar(); e != nil {
						return nil, e
					}
					rn[i2] = c
				}
				u, e = strconv.ParseInt(rn.AsString(), 16, 32)
				if e != nil {
					return nil, e
				}
				c = rune(u)
			default:
				return nil, fmt.Errorf("非法转义字符\\%c", c)
			}
		}
		s = append(s, c)
	}
	return s, nil
}

func (tok *tokenizer) readChar() (rune, error) {
	s, e := tok.readPart(1)
	if e != nil {
		return 0, e
	}
	return s[0], e
}

func (tok *tokenizer) readPart(count int) (JSONString, error) {
	if tok.eof() {
		return nil, errors.New(`data is eof`)
	}
	lastCount := tok.size - tok.offset
	if lastCount < count {
		return nil, errors.New(`last data count less read count`)
	}
	nOffset := tok.offset
	tok.offset += count
	return tok.data[nOffset:tok.offset], nil
}

func (tok *tokenizer) eof() bool {
	if tok == nil {
		return true
	}
	return tok.offset == tok.size
}

func (tok *tokenizer) seek(v int) {
	tok.offset += v
}
