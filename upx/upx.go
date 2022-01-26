package upx

import (
	"fmt"
	"os/exec"
)

func StartUpx() {
	cmd := exec.Command("./upx/upx.exe", "output.exe")
	err := cmd.Run()
	if err == nil {
		fmt.Println("[*] upx success")
	} else {
		fmt.Println("[!] upx error")
	}
}
