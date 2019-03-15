# go-llvm-mc

> Go wrapper for llvm-mc

---

## Dependancies

- [llvm-mc](https://github.com/llvm/llvm-project/tree/master/llvm/tools/llvm-mc)

### Install `llvm-mc`

```bash
$ git clone https://github.com/llvm/llvm-project.git
$ cd llvm-project
$ mkdir build
$ cd build
$ cmake ../llvm -DLLVM_ENABLE_PROJECTS='llvm-objdump,llvm-mc' -DCMAKE_INSTALL_PREFIX=/tmp/llvm
$ make install
$ cp /tmp/llvm/bin/llvm-mc /usr/local/bin/llvm-mc
```

## Getting Started

```bash
$ go run main.go disassemble kernelcache.release.iphone11.decompressed | less
```

```asm
	.section	__TEXT_EXEC,__text,regular,pure_instructions
	mov	x0, #29313
	movk	x0, #28518, lsl #16
	movk	x0, #28786, lsl #32
	movk	x0, #65388, lsl #48
	ret
	neg	w8, w0
	and	w0, w8, #0x7
	ret
	adrp	x8, #22691840
	ldr	x0, [x8]
	ret
	pacibsp
	stp	x24, x23, [sp, #-64]!
	stp	x22, x21, [sp, #16]
	stp	x20, x19, [sp, #32]
	stp	x29, x30, [sp, #48]
	mov	x19, x3
	mov	x20, x2
	mov	x21, x1
	mov	x22, x0
	sub	x23, x5, x4
	mov	x0, x23
	bl	#-68
	mov	w8, #47
	sub	x8, x8, x22
	add	x8, x8, x21
	mov	x9, #-6148914691236517206
	movk	x9, #43691
	umulh	x9, x8, x9
	lsr	x9, x9, #5
        <SNIP>
```

=OR=

Just use `llvm-objdump`

```bash
$ llvm-objdump -arch=arm64 -mattr=v8.5a -d kernelcache.release.iphone11.decompressed | less
```
