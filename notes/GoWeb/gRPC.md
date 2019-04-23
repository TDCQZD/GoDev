## gRPC

### 环境配置
1. go >1.6
2. Install gRPC
```
$ go get -u google.golang.org/grpc
```

``` 
git clone https://github.com/grpc/grpc-go.git $GOPATH/src/google.golang.org/grpc
git clone https://github.com/golang/net.git $GOPATH/src/golang.org/x/net
git clone https://github.com/golang/text.git $GOPATH/src/golang.org/x/text
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
git clone https://github.com/google/go-genproto.git $GOPATH/src/google.golang.org/genproto

cd $GOPATH/src/
go install google.golang.org/grpc

```
3. Install Protocol Buffers v3
> https://github.com/protocolbuffers/protobuf/releases
```
go get -u github.com/golang/protobuf/protoc-gen-go
$ export PATH=$PATH:$GOPATH/bin
```

## 生成gRPC代码
```
protoc --go_out=plugin=protorpc=. arith.proto 
protoc --go_out=. arith.proto [win]
protoc --go_out=grpc:. user.proto

protoc -I=. --cpp_out=. Person.pro，
分别指定目录路径、输出路径和.proto文件路径
```

## ptotorpc 数据序列化与反序列化