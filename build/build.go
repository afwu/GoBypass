package build

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func NormalBuild(code string) {
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

func AdvanceBuild(code string) {
	cmd := []string{
		"/c",
		"go",
		"build",
		"-o",
		"output.exe",
		"-ldflags",
		"-s -w",
		"output/main.go",
	}
	privateBuild(code, cmd)
}

func privateBuild(code string, command []string) {
	_ = os.RemoveAll(filepath.Join(".", "output.exe"))
	newPath := filepath.Join(".", "output")
	_ = os.MkdirAll(newPath, os.ModePerm)
	_ = ioutil.WriteFile("output/main.go", []byte(code), 0777)
	cmd := exec.Command("cmd", command...)
	err := cmd.Run()
	if err == nil {
		fmt.Println("go build success")
	} else {
		fmt.Println(command)
		fmt.Println("go build error")
	}
	_ = os.RemoveAll(newPath)
}
