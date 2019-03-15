# NOTES

```bash
$ echo "mov x1, #10\nretab" | llvm-mc -arch=arm64 -mattr=v8.5a -show-encoding | \
sed 's/.*\(0x[0-9a-f]*\),\(0x[0-9a-f]*\),\(0x[0-9a-f]*\),\(0x[0-9a-f]*\).*'\
'/\1 \2 \3 \4/' | tail -n +2 | llvm-mc -disassemble -arch=arm64 -mattr=v8.5a
```
