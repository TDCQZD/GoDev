# Go 锁
- sync.RWMutex
- sync.Mutex
- sync.WaitGroup 
## sync.RWMutex 读写锁
RWMutex是读写互斥锁。该锁可以被同时多个读取者持有或唯一个写入者持有。

RWMutex可以创建为其他结构体的字段；零值为解锁状态。

RWMutex类型的锁也和线程无关，可以由不同的线程加读取锁/写入和解读取锁/写入锁。



RWMutex方法：

* func (rw *RWMutex) Lock() 
    - Lock方法将rw锁定为写入状态，禁止其他线程读取或者写入

* func (rw *RWMutex) Unlock() 
    -  Unlock方法解除rw的写入锁状态，如果m未加写入锁会导致运行时错误。

* func (rw *RWMutex) RLock() 
    - RLock方法将rw锁定为读取状态，禁止其他线程写入，但不禁止读取。

* func (rw *RWMutex) RUnlock()
    -  Runlock方法解除rw的读取锁状态，如果m未加读取锁会导致运行时错误。

* func (rw *RWMutex) RLocker() Locker
    - Rlocker方法返回一个互斥锁，通过调用rw.Rlock和rw.Runlock实现了Locker接口。
## sync.Mutex 互斥锁
Mutex是一个互斥锁，可以创建为其他结构体的字段；零值为解锁状态。Mutex类型的锁和线程无关，可以由不同的线程加锁和解锁。

互斥锁：同一个时刻只有一个线程能够拿到锁

Mutex方法：
* func (m *Mutex) Lock() 
    -  Lock方法锁住m，如果m已经加锁，则阻塞直到m解锁
* func (m *Mutex) Unlock()
    - Unlock方法解锁m，如果m未加锁会导致运行时错误。锁和线程无关，可以由不同的线程加锁和解锁。
## sync.WaitGroup 
WaitGroup用于等待一组线程的结束。父线程调用Add方法来设定应等待的线程的数量。每个被等待的线程在结束时应调用Done方法。同时，主线程里可以调用Wait方法阻塞至所有线程结束。
* func (wg *WaitGroup) Add(delta int)
    - Add方法向内部计数加上delta，delta可以是负数；如果内部计数器变为0，Wait方法阻塞等待的所有线程都会释放，如果计数器小于0，方法panic。注意Add加上正数的调用应在Wait之前，否则Wait可能只会等待很少的线程。一般来说本方法应在创建新的线程或者其他应等待的事件之前调用。
* func (wg *WaitGroup) Done()
    - Done方法减少WaitGroup计数器的值，应在线程的最后执行。

* func (wg *WaitGroup) Wait()
    - Wait方法阻塞直到WaitGroup计数器减为0。

### 示例：
```
var wg sync.WaitGroup
var urls = []string{
    "http://www.golang.org/",
    "http://www.google.com/",
    "http://www.somestupidname.com/",
}
for _, url := range urls {
    // Increment the WaitGroup counter.
    wg.Add(1)
    // Launch a goroutine to fetch the URL.
    go func(url string) {
        // Decrement the counter when the goroutine completes.
        defer wg.Done()
        // Fetch the URL.
        http.Get(url)
    }(url)
}
// Wait for all HTTP fetches to complete.
wg.Wait()
```
## sync.Once 
```
type Once struct {
    // 包含隐藏或非导出字段
}
```
Once是只执行一次动作的对象。
### 方法
* func (o *Once) Do(f func())
    - Do方法当且仅当第一次被调用时才执行函数f
    - 换句话说，给定变量：`var once Once`如果once.Do(f)被多次调用，只有第一次调用会执行f，即使f每次调用Do 提供的f值不同。需要给每个要执行仅一次的函数都建立一个Once类型的实例。
    - Do用于必须刚好运行一次的初始化。因为f是没有参数的，因此可能需要使用闭包来提供给Do方法调用：`config.once.Do(func() { config.init(filename) })`因为只有f返回后Do方法才会返回，f若引起了Do的调用，会导致死锁。
### 示例：
```
var once sync.Once
onceBody := func() {
    fmt.Println("Only once")
}
done := make(chan bool)
for i := 0; i < 10; i++ {
    go func() {
        once.Do(onceBody)
        done <- true
    }()
}
for i := 0; i < 10; i++ {
    <-done
}
```