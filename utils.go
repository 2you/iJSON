package iJSON

import (
	"strconv"
	"strings"
)

func HashToUInt(s string) (h uint) {
	h = 0
	for i := 0; i < len(s); i++ {
		h = h*129 + uint(s[i]) + 0x9e370001
	}
	return
}

func CompareAvlNode(a, b *AvlNode) int {
	if a.hash < b.hash {
		return -1
	}

	if a.hash > b.hash {
		return +1
	}
	return strings.Compare(a.name, b.name)
}

func CompareName(a, b string) int {
	ha := HashToUInt(a)
	hb := HashToUInt(b)
	if ha < hb {
		return -1
	}

	if ha > hb {
		return +1
	}
	return strings.Compare(a, b)
}

func stripFloatE00(s string) string {
	n := len(s)
	if n < 5 {
		return s
	}

	if s[n-1] == '0' && s[n-2] == '0' && s[n-3] == '+' && s[n-4] == 'e' {
		s = s[:n-4]
	}
	return s
}

func FormatDouble(f float64) (s string) {
	s = strconv.FormatFloat(f, 'e', -1, 64)
	return stripFloatE00(s)
}

func FormatFloat(f float32) (s string) {
	s = strconv.FormatFloat(float64(f), 'e', -1, 32)
	return stripFloatE00(s)
}
