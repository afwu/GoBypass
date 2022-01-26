package build

import (
	"github.com/EmYiQing/GoBypass/log"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func NormalBuild(code string) {
	log.Info("build normal")
	cmd := []string{
		"/c",
		"go",
		"build",
		"-o",
		"output.exe",
		"output/main.go",
	}
	privateBuild(code, cmd)
}

func NormalGarble(code string) {
	log.Info("build normal use garble")
	cmd := []string{
		"build",
		"-o",
		"output.exe",
		"output/main.go",
	}
	privateGrable(code, cmd)
}

func AdvanceBuild(code string) {
	log.Info("build use ldflags")
	cmd := []string{
		"/c",
		"go",
		"build",
		"-o",
		"output.exe",
		"-ldflags",
		"-s -w -H windowsgui",
		"output/main.go",
	}
	privateBuild(code, cmd)
}

func AdvanceGarble(code string) {
	log.Info("build use ldflags and garble")
	cmd := []string{
		"build",
		"-o",
		"output.exe",
		"-ldflags",
		"-s -w -H windowsgui",
		"output/main.go",
	}
	privateGrable(code, cmd)
}

func privateGrable(code string, command []string) {
	_ = os.RemoveAll(filepath.Join(".", "output.exe"))
	newPath := filepath.Join(".", "output")
	_ = os.MkdirAll(newPath, os.ModePerm)
	_ = ioutil.WriteFile("output/main.go", []byte(code), 0777)
	cmd := exec.Command("./garble/garble.exe", command...)
	err := cmd.Run()
	if err == nil {
		log.Info("build success")
		log.Info("file: output.exe")
	} else {
		log.Error("error")
	}
	_ = os.RemoveAll(newPath)
}

func privateBuild(code string, command []string) {
	_ = os.RemoveAll(filepath.Join(".", "output.exe"))
	newPath := filepath.Join(".", "output")
	_ = os.MkdirAll(newPath, os.ModePerm)
	_ = ioutil.WriteFile("output/main.go", []byte(code), 0777)
	cmd := exec.Command("cmd", command...)
	err := cmd.Run()
	if err == nil {
		log.Info("build success")
		log.Info("file: output.exe")
	} else {
		log.Error("error")
	}
	_ = os.RemoveAll(newPath)
}
