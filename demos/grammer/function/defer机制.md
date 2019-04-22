# Defer

函数中，程序员经常需要创建资源(比如：数据库连接、文件句柄、锁等) ，为了在函数执行完毕后，及时的释放资源，Go的设计者提供defer (延时机制)。

```
func sum(n1 int, n2 int) int {

    defer fmt.Println("ok1 n1=", n1) // sum defer栈  
    defer fmt.Println("ok2 n2=", n2) // sum defer栈   
    res := n1 + n2 // res = 32
    fmt.Println("ok3 res=", res) // (1) ok3 res = 32
    return res
}
func main() {
    res := sum(10, 20) 
}

```

1. 当go执行到一个defer时，不会立即执行defer后的语句，而是将defer 后的语句压入到一个栈中[可以称该栈为defer栈], 然后继续执行函数下一个语句。 
2. 当函数执行完毕后，在从defer栈中，依次从栈顶取出语句执行(注：遵守栈 先入后出的机制)。 
3. 在defer 将语句放入到栈时，也会将相关的值拷贝同时入栈。 

