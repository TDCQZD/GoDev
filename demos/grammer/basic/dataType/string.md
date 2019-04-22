## 字符串
Go中的字符串都是采用UTF-8字符集编码。字符串是用一对双引号（""）或反引号（``）括起来定义，它的类型是string。

字符串就是一串固定长度的字符连接起来的字符序列。Go的字符串是由单个字节连接起来的。 也就是说对于传统的字符串是由字符组成的，而Go的字符串不同，它是由字节组成的。

```
//示例代码
var frenchHello string  // 声明变量为字符串的一般方法
var emptyString string = ""  // 声明了一个字符串变量，初始化为空字符串
func test() {
    no, yes, maybe := "no", "yes", "maybe"  // 简短声明，同时声明多个变量
    japaneseHello := "Konichiwa"  // 同上
    frenchHello = "Bonjour"  // 常规赋值
}
```
## 字符串操作

### 修改字符串
在Go中字符串是不可变的，例如下面的代码编译时会报错：`cannot assign to s[0]`
```
var s string = "hello"
s[0] = 'c'
```
但如果真的想要修改怎么办呢？下面的代码可以实现：
```
s := "hello"
c := []byte(s)  // 将字符串 s 转换为 []byte 类型
c[0] = 'c'
s2 := string(c)  // 再转换回 string 类型
fmt.Printf("%s\n", s2)
```
### + 拼接

Go中可以使用+操作符来连接两个字符串：
```
s := "hello,"
m := " world"
a := s + m
fmt.Printf("%s\n", a)
```
修改字符串也可写为：
```
s := "hello"
s = "c" + s[1:] // 字符串虽不能更改，但可进行切片操作
fmt.Printf("%s\n", s)
```
> 当一行字符串太长时，需要使用到多行字符串时，可以使用+拼接 .
> 注意：+要放在字符串后面，放在前面会报错。invalid operation: + untyped string
```
	ss:= "sdfghjk"
	+"weryuio"
```
### 声明多行的字符串
通过``来声明：
```
m := `hello
    world`
``` 
