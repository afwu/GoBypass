package build

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func NormalBuild(code string) {
	fmt.Println("[*] build normal")
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
	fmt.Println("[*] build normal use garble")
	cmd := []string{
		"build",
		"-o",
		"output.exe",
		"output/main.go",
	}
	privateGrable(code, cmd)
}

func AdvanceBuild(code string) {
	fmt.Println("[*] build use ldflags")
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
	fmt.Println("[*] build use ldflags and garble")
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
		fmt.Println("[*] build success")
		fmt.Println("[*] file: output.exe")
	} else {
		fmt.Println("[!] error")
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
		fmt.Println("[*] build success")
		fmt.Println("[*] file: output.exe")
	} else {
		fmt.Println("[!] error")
	}
	_ = os.RemoveAll(newPath)
}
