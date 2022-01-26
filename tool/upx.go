package tool

import (
	"github.com/EmYiQing/GoBypass/log"
	"os/exec"
)

func StartUpx() {
	cmd := exec.Command("./tool/upx.exe", "output.exe")
	err := cmd.Run()
	if err == nil {
		log.Info("upx success")
	} else {
		log.Error("upx error")
	}
}
