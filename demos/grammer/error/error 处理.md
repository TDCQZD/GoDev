# error类型
Go内置有一个error类型，专门用来处理错误信息，Go的package里面还专门有一个包errors来处理错误：
```
err := errors.New("emit macho dwarf: elf header corrupted")
if err != nil {
	fmt.Print(err)
```