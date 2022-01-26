package encode

import (
	"strconv"
)

var xorKey = []byte{0x13, 0x54, 077, 0x1A, 0xA1, 0x3F, 0x04, 0x8B}

func Encode(src string) string {
	var result string
	j := 0
	s := ""
	bt := []rune(src)
	for i := 0; i < len(bt); i++ {
		s = strconv.FormatInt(int64(byte(bt[i])^xorKey[j]), 16)
		if len(s) == 1 {
			s = "0" + s
		}
		result = result + (s)
		j = (j + 1) % 8
	}
	return result
}

func Decode(src string) string {
	var result string
	var s int64
	j := 0
	bt := []rune(src)
	for i := 0; i < len(src)/2; i++ {
		s, _ = strconv.ParseInt(string(bt[i*2:i*2+2]), 16, 0)
		result = result + string(byte(s)^xorKey[j])
		j = (j + 1) % 8
	}
	return result
}
