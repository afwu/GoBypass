package parser

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

func ParseShellCode() string {
	fmt.Println("[*] parse shellcode")
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

func GetFinalCode(module string, shellcode string) string {
	template, _ := ioutil.ReadFile(fmt.Sprintf("core/%s/%s.go", module, module))
	return strings.ReplaceAll(string(template), "__SHELLCODE__", shellcode)
}
