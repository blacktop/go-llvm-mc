# NOTES

```bash
$ echo "mov x1, #10\nretab" | llvm-mc -arch=arm64 -mattr=v8.5a -show-encoding | \
sed 's/.*\(0x[0-9a-f]*\),\(0x[0-9a-f]*\),\(0x[0-9a-f]*\),\(0x[0-9a-f]*\).*'\
'/\1 \2 \3 \4/' | tail -n +2 | llvm-mc -disassemble -arch=arm64 -mattr=v8.5a
```

```bash
$ echo "mov x1, #10\nret" | llvm-mc -arch=arm64 -mattr=v8.5a -show-encoding | \
sed 's/.*0x\([0-9a-f]*\),0x\([0-9a-f]*\),0x\([0-9a-f]*\),0x\([0-9a-f]*\).*'\
'/\1 \2 \3 \4/' | tail -n +2 | sed ':a;N;$!ba;s/\n/ /g'
```

```bash
$ cstool arm64 "41 01 80 d2 c0 03 5f d6"

 0  41 01 80 d2  movz   x1, #0xa
 4  c0 03 5f d6  ret
```

```bash
$ echo "0xd6 0x30 0xf0 0xd5" | /usr/local/opt/llvm/bin/llvm-mc -disassemble -arch=arm64 -mattr=v8.5a 
```
