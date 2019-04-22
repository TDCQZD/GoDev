# 使用 jsoniter 构建

高性能100％兼容直接替代“encoding / json”.
> 来源滴滴出行

## 依赖
```
go get github.com/json-iterator/go
```
## 使用
### 示例一:Marshal
Replace
```
import "encoding/json"
json.Marshal(&data)
```
with

```
import "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary
json.Marshal(&data)
```
### 示例二:Unmarshal
Replace
```
import "encoding/json"
json.Unmarshal(input, &data)
```
with
```
import "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary
json.Unmarshal(input, &data)
```