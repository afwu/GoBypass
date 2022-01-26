package log

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func Info(format string, a ...interface{}) {
	now := getTime()
	var data string
	if checkOS() {
		data = fmt.Sprintf("[+] [%s] %s\n", now, format)
	} else {
		data = fmt.Sprintf("\x1b[32m[+] [%s] %s\x1b[0m\n", now, format)
	}
	_, _ = fmt.Fprintf(os.Stdout, data, a...)
}

func Error(format string, a ...interface{}) {
	now := getTime()
	var data string
	if checkOS() {
		data = fmt.Sprintf("[-] [%s] %s\n", now, format)
	} else {
		data = fmt.Sprintf("\x1b[31m[-] [%s] %s\x1b[0m\n", now, format)
	}
	_, _ = fmt.Fprintf(os.Stdout, data, a...)
}

func getTime() string {
	currentTime := time.Now().Format("15:04:05")
	return currentTime
}

func checkOS() bool {
	sysType := runtime.GOOS
	if sysType == "windows" {
		return true
	}
	return false
}
