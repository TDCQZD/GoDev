# random number
## 1、math/rand包 数据可能重复
```
func main() {
    fmt.Println("My first lucky number is", rand.Intn(10))
    fmt.Println("My senond lucky number is", rand.Intn(10))
}
```
此程序的运行环境是固定的，因此 rand.Intn 总是会返回相同的数字

```
func main() {
    rand.Seed(time.Now().Unix())
    rand.Seed(time.Now().UnixNano())
    fmt.Println("My first lucky number is", rand.Intn(10))
    fmt.Println("My senond lucky number is", rand.Intn(10))
}
```
设置随机数种子，可以保证每次随机都是随机的

## 2、不可重复随机数
```
rad := rand.New(rand.NewSource(time.Now().Unix()))
for i := 0; i < rad.Intn(9)+1; i++ {
fmt.Println(rad.Intn(50))
}
```