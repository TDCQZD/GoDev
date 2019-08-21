# golang中死锁的情况分析

## channel
### 1、直接读取空channel的死锁
当一个channel中没有数据，而直接读取时，会发生死锁：
```
func main() {
    q := make(chan int, 2)
    <-q
}
```
**解决方法：**

使用select方法阻止，在default中放置默认处理方式：
```
func main() {
    q := make(chan int, 2)
    select {
    case v := <-q:
        fmt.Println(v)
    default:
        fmt.Println("nothing in channel")
    }
}
```
### 2、过度写入数据造成的死锁
写入数据超过channel的容量，也会造成死锁：
```
func main() {
    q := make(chan int, 2)
    q <- 1
    q <- 2
    q <- 3
}
```
**解决方案：**

与写入的方式一样，使用select方法阻止，在default中放置默认处理方式：
```
func main() {
    q := make(chan int, 2)
    q <- 1
    q <- 2
    select {
    case q <- 3:
        fmt.Println("ok")
    default:
        fmt.Println("wrong")
    }

}
```

### 3、向已经关闭的channel中写入数据
这种造成的不是死锁，而是产生panic。
```
func main() {
    q := make(chan int, 2)
    close(q)
    q <- 1
}
```
**解决方案**：

只有别向已经关闭的channel写数据。

## goroutine
正确示例：
```
func main() {
	c:=make(chan string)
	go func() {
		c<-"hello"
	}()
	fmt.Println(<-c)
}
```
错误示例：
```
package main

import "fmt"

func main() {
	c:=make(chan string)
	fmt.Println(<-c)
	go func() {
		c<-"hello"
	}()
}
```
上面这个是错的，会发生死锁，因为程序会阻塞在`fmt.Println(<-c)`，并不会向下执行。在该行发生了阻塞，以致于下面的代码没有机会执行。因此可以总结写在读前，写操作完成后才有读操作。

## 资源竞争发生死锁
```
type value struct {
    mu    sync.Mutex
    value int
}

func main() {
    var wg sync.WaitGroup

    printSum := func(v1, v2 *value) {
        defer wg.Done()
        v1.mu.Lock()
        defer v1.mu.Unlock()
        time.Sleep(2 * time.Second)
        v2.mu.Lock()
        defer v2.mu.Unlock()
        fmt.Printf("Sum=%v\n", v1.value+v2.value)
    }

    var a, b value
    wg.Add(2)
    go printSum(&a, &b)
    go printSum(&b, &a)
    wg.Wait()
}
```
a，b互相等待对方释放锁，形成死锁