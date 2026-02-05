# hello world

## 1. 根据 proto 文件生成 go 文件

### 下载protoc

```shell
cd ~/dvp/dvp_go/bin
wget https://github.com/protocolbuffers/protobuf/releases/download/v33.1/protoc-33.1-linux-x86_64.zip
```

### 安装ptotoc-gen-go(安装新的)

```shell
https://github.com/golang/protobuf/tree/master/protoc-gen-go 这是旧的
https://github.com/protocolbuffers/protobuf-go/tree/master/cmd/protoc-gen-go 这是新的

# 安装旧的 protoc-gen-go (我打的proto文件，必须用旧版本的protoc-gen-go)
go install github.com/golang/protobuf/protoc-gen-go@v1.3.5

# 安装新的protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

### protoc-gen-micro

```
go install go-micro.dev/v5/cmd/protoc-gen-micro@v5.10.0

or

# 手工编译
cd /home/robert/dvp/gitee/go-micro/cmd/protoc-gen-micro
go build
cp ./protoc-gen-micro /home/robert/dvp/dvp_go/bin
```

## 2, 编译

```
根据proto文件生成go文件
./gen_pb.sh

cd GreeterServer
go build

cd gateway
go build

```

服务发现

- 默认使用 MDNS
- consul

```

curl -X POST http://localhost:9098/api/GreeterServer/SayHello \
 -H "Content-Type: application/json" \
 -d '{"name":"robert"}'

```
