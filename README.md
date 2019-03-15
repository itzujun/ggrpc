
# grpc简单demo

```
项目结构
├──consul|    #接口定义文件
         | grpcfile.pb.go # 命令生成接口文件
         | grpcfile.proto # 接口自定义文件
├── client.go  # grpc客户端
├── server.go  # grpc服务端
```


### 项目构建
1. 安装golang.org包 (参考：https://www.jianshu.com/p/096c5c253f75)
2. 新建接口 grpcfile.proto
3. 在grpcfile.proto目录下执行protoc --go_out=plugins=grpc:.  .\grpcfile.proto 生成grpcfile.pb.go接口文件
4. 建立服务端server.go
6. 建立客户端client.go

```
end...
```













